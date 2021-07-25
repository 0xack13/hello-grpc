package main

import (
	"context"
	"fmt"
	"log"

	"github.com/0xack13/hello/hellopb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello client..")
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("Error in %v", err)
	}
	defer cc.Close()
	client := hellopb.NewHelloServiceClient(cc)
	request := &hellopb.HelloRequest{Name: "Jimmy"}
	resp, _ := client.Hello(context.Background(), request)
	fmt.Println(resp.Greeting)
}
