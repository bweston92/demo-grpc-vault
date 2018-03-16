package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/bweston92/demo-grpc-vault/pb/hello"
	vault "github.com/hashicorp/vault/api"
	"google.golang.org/grpc"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
)

func main() {
	ssi := grpc_middleware.ChainStreamServer(
		grpc_ctxtags.StreamServerInterceptor(),
		grpc_prometheus.StreamServerInterceptor,
		grpc_recovery.StreamServerInterceptor(),
	)
	usi := grpc_middleware.ChainUnaryServer(
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_prometheus.UnaryServerInterceptor,
		grpc_recovery.UnaryServerInterceptor(),
	)

	srv := grpc.NewServer(
		grpc.StreamInterceptor(ssi),
		grpc.UnaryInterceptor(usi),
	)

	// Register the implementation for hello server.
	hello.RegisterHelloServer(srv, getHelloImplementation())

	nl, err := makeGRPCListener()
	if err != nil {
		log.Fatalf("unable to get gRPC listener: %s", err)
	}

	log.Fatal(srv.Serve(nl))
}

/////////////////////////////////////////////////////////////
// The following lines contain the the logic for getting a
// certificate and listener with the Vault token provided.
/////////////////////////////////////////////////////////////

func makeGRPCListener() (net.Listener, error) {
	v, err := vault.NewClient(vault.DefaultConfig())
	if err != nil {
		return nil, err
	}

	if m, err := v.Sys().ListMounts(); err != nil {
		return nil, err
	} else if _, ok := m["pki/"]; !ok {
		return nil, errors.New("Vault needs PKI secret mounted")
	}

	// @todo get certificate from Vault (v)
	return net.Listen("tcp", ":8000")
}

/////////////////////////////////////////////////////////////
// The following lines contain the simple implementation for
// the gRPC service "Hello"
/////////////////////////////////////////////////////////////

// Simple is just an "simple" implementation of hello.
type Simple struct{}

func getHelloImplementation() hello.HelloServer {
	return &Simple{}
}

// SayHelloToName will take the request and extract the name
// returning the name in a hello world sting.
func (s *Simple) SayHelloToName(_ context.Context, req *hello.SayHelloToNameRequest) (*hello.SayHelloToNameResponse, error) {
	return &hello.SayHelloToNameResponse{
		Text: fmt.Sprintf("Hello %s, welcome to demo-grpc-vault!", req.GetFullName()),
	}, nil
}
