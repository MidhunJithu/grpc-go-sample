package main

import (
	"context"
	"log"

	pb "github.com/MidhunJithu/grpc-go-sample/blog/proto"
	"github.com/MidhunJithu/grpc-go-sample/blog/server/models"
	"github.com/MidhunJithu/grpc-go-sample/blog/server/repo"
)

type blogServer struct {
	pb.UnimplementedBlogServiceServer
	repo repo.Blog
}

func NewBlogServer(repo repo.Blog) *blogServer {
	return &blogServer{
		repo: repo,
	}
}

func (s *blogServer) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Print("recievd create request")
	blogItem, err := models.ParseBlog(in)
	if err != nil {
		log.Printf("failed to parse blog %v", err)
		return nil, err
	}
	err = s.repo.Create(ctx, blogItem)
	if err != nil {
		log.Printf("failed to create blog %v", err)
		return nil, err
	}
	return &pb.BlogId{Id: blogItem.ID.Hex()}, nil
}
