package main

import (
	"log"
	"net"

	pb "github.com/MidhunJithu/grpc-go-sample/blog/proto"
	"github.com/MidhunJithu/grpc-go-sample/blog/server/repo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func main() {

	list, err := net.Listen("tcp", "0.0.0.0.:50051")
	if err != nil {
		log.Fatalf("failed to listne to tcp traffic %v", err)
	}

	// ssl auth
	certFile := "ssl/server.crt"
	keyFile := "ssl/server.pem"
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("failed loading certificates: %v", err)
	}
	// create grpc server
	server := grpc.NewServer(grpc.Creds(creds))

	// dependedncies
	blogRepo := repo.NewBlogRepo("blogdb", "blogs", "mongodb://root:mypassword@localhost:27017/")
	blogSrv := NewBlogServer(blogRepo)

	// register service
	pb.RegisterBlogServiceServer(server, blogSrv)

	// reflection
	reflection.Register(server)

	log.Printf("listening on %s\n", list.Addr())
	// serve
	if err = server.Serve(list); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
