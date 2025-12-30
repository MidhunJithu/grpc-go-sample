package main

import (
	"log"

	"github.com/MidhunJithu/grpc-go-sample/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addess = "0.0.0.0:50051"

func main() {
	conn, err := grpc.NewClient(
		addess,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("failed to dail in to server at %v, %v", addess, err)
	}
	defer conn.Close()
	client := proto.NewCalculatorServiceClient(conn)
	doSum(client)
	getPrimeFactors(client)
}
