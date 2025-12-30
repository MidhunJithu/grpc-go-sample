package main

import (
	"context"
	"log"

	"github.com/MidhunJithu/grpc-go-sample/calculator/proto"
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
