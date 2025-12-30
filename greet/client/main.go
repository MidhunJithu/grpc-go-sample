package main

import (
	"log"

	pb "github.com/MidhunJithu/grpc-go-sample/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var address = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to grpc server %v", err)
	}
	log.Print("connected to grpc server")

	client := pb.NewGreetServiceClient(conn)

	doGreet(client)
	defer conn.Close()
}
