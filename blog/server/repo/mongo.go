package repo

import (
	"context"
	"log"

	"github.com/MidhunJithu/grpc-go-sample/blog/server/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogImpl struct {
	collection string
	DB         *mongo.Client
}

func NewBlogRepo(collection string, URI string) Blog {
	opts := options.Client().ApplyURI(URI)
	ctx := context.Background()
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("failed to load monogo client %v", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("failed to ping monogo client %v", err)
	}
	return &BlogImpl{
		collection: collection,
		DB:         client,
	}
}

// Create implements [Blog].
func (b *BlogImpl) Create(context.Context, *models.Blog) error {
	log.Print("Implement me")
	return nil
}
