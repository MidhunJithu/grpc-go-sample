package main

import (
	"context"
	"log"

	"github.com/MidhunJithu/grpc-go-sample/calculator/proto"
)

func doSum(client proto.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	req := proto.SumRequest{
		FirstNumber:  1111,
		SecondNumber: 5316,
	}
	res, err := client.Sum(context.Background(), &req)
	if err != nil {
		log.Fatalf("failed to invoke Sum %v", err)
	}

	log.Printf("Sum of %v and %v is %v", req.FirstNumber, req.SecondNumber, res.Result)
}
