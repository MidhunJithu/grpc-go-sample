package main

import (
	"log"
	"net"

	pb "github.com/MidhunJithu/grpc-go-sample/greet/proto"
	"google.golang.org/grpc"
)

var serverAdress = "0.0.0.0:50051"

func main() {

	lis, err := net.Listen("tcp", serverAdress)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
	log.Printf("Listening on %s\n", serverAdress)
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

	log.Printf("Server started on %s\n", serverAdress)
}
