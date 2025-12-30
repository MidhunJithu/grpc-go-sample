package main

import (
	"context"
	"crypto/rand"
	"io"
	"log"
	"math/big"

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

func getPrimeFactors(client proto.CalculatorServiceClient) {
	log.Println("getPrimeFactors was invoked")
	number, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		log.Fatalf("cannot create the number for prime factorisation")
	}

	log.Printf("calling prime composition with %d", number.Int64())
	resp, err := client.PrimeComposition(context.Background(),
		&proto.PrimesRequest{
			Number: int32(number.Int64()),
		},
	)

	for {
		out, err := resp.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to get the prime factors %v", err)
		}
		log.Printf("prime factor %v", out.Result)
	}
}
