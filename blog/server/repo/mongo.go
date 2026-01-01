package repo

import (
	"context"
	"log"

	"github.com/MidhunJithu/grpc-go-sample/blog/server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogImpl struct {
	collection *mongo.Collection
	client     *mongo.Client
}

func NewBlogRepo(db string, collection string, URI string) Blog {
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
		collection: client.Database(db).Collection(collection),
		client:     client,
	}
}

// Create implements [Blog].
func (b *BlogImpl) Create(ctx context.Context, req *models.Blog) error {
	res, err := b.collection.InsertOne(ctx, req)
	if err != nil {
		return err
	}
	iD, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Printf("failed to get inserted data %v", res)
		return mongo.ErrUnacknowledgedWrite
	}
	req.ID = iD
	return nil
}
