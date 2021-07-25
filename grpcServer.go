package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/0xack13/hello/hellopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
}

func (*server) Hello(ctx context.Context, request *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	name := request.Name
	response := &hellopb.HelloResponse{
		Greeting: "Hello " + name,
	}
	return response, nil
}

func main() {
	address := "0.0.0.0:50051"
	lis, erro := net.Listen("tcp", address)
	if erro != nil {
		log.Fatalf("Error %v", erro)
	}
	fmt.Printf("Server is listening on %v...", address)
	s := grpc.NewServer()
	hellopb.RegisterHelloServiceServer(s, &server{})
	reflection.Register(s)
	s.Serve(lis)
}
