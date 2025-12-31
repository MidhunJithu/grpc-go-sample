package main

import (
	"context"
	"io"
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

func (c *calculator) GetAverages(req grpc.ClientStreamingServer[proto.AvgRequest, proto.AvgResponse]) error {

	log.Printf("recieved the Getavaerages call from the grpc client")

	sum := int32(0)
	count := 0
	for {
		msg, err := req.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		sum += msg.Number
		count++
	}
	avg := float32(sum) / float32(count)

	return req.SendAndClose(&proto.AvgResponse{
		Result: avg,
	})
}
