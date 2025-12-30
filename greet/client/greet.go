package main

import (
	"context"
	"io"
	"log"

	pb "github.com/MidhunJithu/grpc-go-sample/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Printf("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Midhun Jithu",
	})
	if err != nil {
		log.Fatalf("could not greet: %v\n", err)
	}
	log.Printf("Greeting: %s\n", res.Result)
}

func doGreetmany(client pb.GreetServiceClient) {

	res, err := client.GreetMany(context.Background(), &pb.GreetRequest{
		FirstName: "Midhun Jithu",
	})
	if err != nil {
		log.Fatalf("failed to make server streaming grpc call %v", err)
	}

	for {
		out, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to recive streaming data %v", err)
		}
		log.Printf("Recieved streaming data %v", out.Result)
	}
}

func doLongGreet(c pb.GreetServiceClient) {
	log.Printf("doLongGreet was invoked")
	reqs := []*pb.GreetRequest{
		{FirstName: "Midhun"},
		{FirstName: "Jithu"},
		{FirstName: "Gayu"},
	}
	stream, err := c.LongGreets(context.Background())
	if err != nil {
		log.Fatalf("failed to call client streaming rpc %v", err)
	}

	for _, req := range reqs {
		if err := stream.Send(req); err != nil {
			log.Fatalf("failed to send request to client streaming rpc %v", err)
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to recieve response from client streaming call %v", err)
	}

	log.Printf("Recieved response from client streaming call %v", resp.Result)
}
