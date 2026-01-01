package main

import (
	"log"
	"net"

	pb "github.com/MidhunJithu/grpc-go-sample/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var serverAdress = "0.0.0.0:50051"

func main() {

	lis, err := net.Listen("tcp", serverAdress)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}
	log.Printf("Listening on %s\n", serverAdress)
	tlsCred, err := credentials.NewServerTLSFromFile("ssl/server.crt", "ssl/server.pem")
	if err != nil {
		log.Fatalf("failed to generate tls from file %v", err)
	}
	s := grpc.NewServer(grpc.Creds(tlsCred))
	pb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

	log.Printf("Server started on %s\n", serverAdress)
}
