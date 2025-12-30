package main

import (
	"context"
	"log"

	"github.com/MidhunJithu/grpc-go-sample/calculator/proto"
	"google.golang.org/grpc"
)

type calculator struct {
	proto.UnimplementedCalculatorServiceServer
}

func (c *calculator) Sum(ctx context.Context, in *proto.SumRequest) (*proto.SumResponse, error) {
	log.Printf("recived the sum call from the grpc client %v", in)
	return &proto.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}

func (c *calculator) PrimeComposition(req *proto.PrimesRequest, stream grpc.ServerStreamingServer[proto.PrimesResponse]) error {
	log.Printf("recived the PrimeComposition call from the grpc client %v", req)
	number := req.Number
	divisor := int32(2)

	for number > 1 {
		if number%divisor == 0 {
			number = number / divisor
			stream.Send(&proto.PrimesResponse{
				Result: divisor,
			})
			continue
		}
		if divisor == 2 {
			divisor++
			continue
		}
		divisor += 2
	}
	return nil
}
