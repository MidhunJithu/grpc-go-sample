package main

import (
	"log"
	"net"

	"github.com/MidhunJithu/grpc-go-sample/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var address = "0.0.0.0:50051"

func main() {
	list, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listent to %v, %v", address, err)
	}
	log.Printf("Listening on %s\n", address)

	server := grpc.NewServer()
	proto.RegisterCalculatorServiceServer(server, &calculator{})

	reflection.Register(server)
	if err := server.Serve(list); err != nil {
		log.Fatalf("failed to server on the grpc addess %v", err)
	}
	log.Printf("Server started on %s\n", address)
}
