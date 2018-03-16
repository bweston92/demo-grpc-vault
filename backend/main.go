package main

import (
	"log"
	"net"
	"os"

	"github.com/bweston92/demo-grpc-vault/pb/hello"
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

func makeGRPCListener() (net.Listener, error) {
	_ = os.Getenv("VAULT_ADDR")
	_ = os.Getenv("VAULT_TOKEN")

	// @todo get certificate from Vault

	return net.Listen("tcp", ":8000")
}
