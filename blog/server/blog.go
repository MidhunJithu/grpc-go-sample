package main

import (
	"context"
	"errors"
	"log"

	pb "github.com/MidhunJithu/grpc-go-sample/blog/proto"
	"github.com/MidhunJithu/grpc-go-sample/blog/server/models"
	"github.com/MidhunJithu/grpc-go-sample/blog/server/repo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *blogServer) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.BlogResponse, error) {
	log.Print("recievd read request")
	blogItem, err := s.repo.Read(ctx, in.Id)
	if err != nil {
		log.Printf("failed to read blog %v", err)
		return nil, err
	}
	return models.ParseBlogItem(blogItem), nil
}

func (s *blogServer) ListBlogs(filter *pb.ListFilter, stream grpc.ServerStreamingServer[pb.BlogResponse]) error {
	log.Print("recievd list request")
	ctx := context.Background()

	// set deafult filters
	filterData := models.ParseFilter(filter)

	blogs, err := s.repo.List(ctx, *filterData)
	if err != nil {
		log.Printf("failed to list blogs %v", err)
		return status.Error(codes.Internal, "failed to list the items")
	}

	// stream the blogs to the client
	for _, blog := range blogs {
		blogItem := models.ParseBlogItem(blog)
		if err = stream.Send(blogItem); err != nil {
			log.Printf("failed to send blog %v", err)
			return status.Error(codes.Internal, "failed to send the items")
		}
	}

	return nil
}

func (s *blogServer) UpdateBlog(ctx context.Context, in *pb.UpdateBlogRequest) (*pb.BlogResponse, error) {
	log.Print("recievd update request")
	blogItem, err := models.ParseUpdaeBlog(in)
	if err != nil {
		log.Printf("failed to parse blog  input %v", err)
		return nil, status.Error(codes.InvalidArgument, "Failed to parse input data, please check the input")
	}
	blogItem, err = s.repo.Update(ctx, blogItem)
	if err != nil {
		if errors.Is(err, repo.ErrBlogNotFound) {
			return nil, status.Error(codes.NotFound, "blog not found")
		}
		log.Printf("failed to update blog %v", err)
		return nil, status.Error(codes.Internal, "failed to update the field")
	}
	return models.ParseBlogItem(blogItem), nil
}
