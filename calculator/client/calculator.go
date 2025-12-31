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

	factors := make([]int32, 0)
	for {
		out, err := resp.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to get the prime factors %v", err)
		}
		factors = append(factors, out.Result)
	}
	log.Printf("prime factor of %v is %v", number.Int64(), factors)
}

func getAverage(client proto.CalculatorServiceClient) {
	log.Println("getAverage was invoked")
	stream, err := client.GetAverages(context.Background())
	if err != nil {
		log.Fatalf("failed to invoke the grpc client streaming call")
	}
	numbers := make([]int32, 0, 10)
	for range 10 {
		number, err := rand.Int(rand.Reader, big.NewInt(100))
		if err != nil {
			log.Fatalf("failed to generate number for averages function %v", err)
		}
		numbers = append(numbers, int32(number.Int64()))
		err = stream.Send(&proto.AvgRequest{
			Number: int32(number.Int64()),
		})
		if err != nil {
			log.Fatalf("failed to send the number for averages %v", err)
		}

	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to recieve the avaerage from the grpc server %v", err)
	}
	log.Printf("the average of %v  is %v", numbers, resp.Result)
}
