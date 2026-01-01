package main

import (
	"log"
	"time"

	pb "github.com/MidhunJithu/grpc-go-sample/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var address = "localhost:50051"

func main() {
	tlsCred, err := credentials.NewClientTLSFromFile("ssl/ca.crt", "")
	if err != nil {
		log.Fatalf("failed to create creds from tls file %v", err)
	}
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(tlsCred))
	if err != nil {
		log.Fatalf("failed to connect to grpc server %v", err)
	}
	defer conn.Close()

	log.Print("connected to grpc server")

	client := pb.NewGreetServiceClient(conn)

	doGreet(client)
	doGreetmany(client)
	doLongGreet(client)
	doGreetEveryone(client)
	doGreetWithDeadline(client, 2*time.Second)
}
