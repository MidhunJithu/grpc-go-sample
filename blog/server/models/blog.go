package models

import (
	"time"

	"github.com/MidhunJithu/grpc-go-sample/blog/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const maxFilterLimit = 150

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

type Filter struct {
	AuthorID     string
	Title        string
	CreatedAtGTE time.Time
	Pagination
}

type Pagination struct {
	Limit  int32
	Offset int32
	Page   int32
}

func ParseFilter(filter *proto.ListFilter) *Filter {
	filterItem := &Filter{}
	filterItem.AuthorID = filter.AuthorId
	filterItem.Title = filter.Title
	filterItem.CreatedAtGTE = filter.CreatedAtGte.AsTime()
	filterItem.Limit = filter.Limit
	filterItem.Page = filter.Page

	if filterItem.Page <= 0 {
		filterItem.Page = 1
	}

	if filterItem.Limit <= 0 {
		filterItem.Limit = 10
	}
	if filterItem.Limit > maxFilterLimit {
		filterItem.Limit = maxFilterLimit
	}
	filterItem.Pagination.Offset = (filterItem.Page - 1) * filterItem.Pagination.Limit

	return filterItem
}

func ParseUpdaeBlog(in *proto.UpdateBlogRequest) (*Blog, error) {

	id, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, err
	}
	blogItem := &Blog{
		ID:      id,
		Title:   in.Title,
		Content: in.Content,
	}
	return blogItem, nil
}
