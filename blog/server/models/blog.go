package models

import (
	"time"

	"github.com/MidhunJithu/grpc-go-sample/blog/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID  string             `bson:"author_id"`
	Title     string             `bson:"title"`
	Content   string             `bson:"content"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func ParseBlog(b *proto.Blog) (*Blog, error) {
	blogItem := &Blog{}
	blogItem.AuthorID = b.AuthorId
	blogItem.Title = b.Title
	blogItem.Content = b.Content
	return blogItem, nil
}

func ParseBlogItem(blogItem *Blog) *proto.BlogResponse {
	return &proto.BlogResponse{
		Id:        blogItem.ID.Hex(),
		AuthorId:  blogItem.AuthorID,
		Title:     blogItem.Title,
		Content:   blogItem.Content,
		CreatedAt: timestamppb.New(blogItem.CreatedAt),
		UpdatedAt: timestamppb.New(blogItem.UpdatedAt),
	}
}
