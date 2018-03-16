package main

import (
	"context"
	"fmt"

	"github.com/bweston92/demo-grpc-vault/pb/hello"
)

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
