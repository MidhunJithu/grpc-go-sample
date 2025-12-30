package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"math/big"
	"strings"

	pb "github.com/MidhunJithu/grpc-go-sample/greet/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreetServiceServer
}

func (s *server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Recieved unary call with params %v", in)

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}

func (s *server) GreetMany(in *pb.GreetRequest, stream grpc.ServerStreamingServer[pb.GreetResponse]) error {
	log.Printf("Recieved streaming call with params %v", in)

	n, err := rand.Int(rand.Reader, big.NewInt(50))
	if err != nil {
		return err
	}
	max := n.Int64()
	for i := range max {
		response := &pb.GreetResponse{
			Result: fmt.Sprintf("Hello %s Number %d", in.FirstName, i),
		}
		err = stream.Send(response)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *server) LongGreets(req grpc.ClientStreamingServer[pb.GreetRequest, pb.GreetResponse]) error {

	log.Printf("Recieved streaming call with params %v", req)
	var res strings.Builder

	for {
		msg, err := req.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to recieve client streaming call %v", err)
		}
		res.WriteString("Hello " + msg.FirstName + "! " + "\n")
	}
	req.SendAndClose(&pb.GreetResponse{
		Result: res.String(),
	})
	return nil
}
