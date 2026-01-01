package repo

import (
	"context"
	"log"
	"time"

	"github.com/MidhunJithu/grpc-go-sample/blog/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogImpl struct {
	collection *mongo.Collection
	client     *mongo.Client
}

var _ Blog = (*BlogImpl)(nil)

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
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()
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

func (b *BlogImpl) Read(ctx context.Context, id string) (*models.Blog, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var blogItem models.Blog
	err = b.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&blogItem)
	if err != nil {
		return nil, err
	}
	return &blogItem, nil
}

// List implements [Blog].
func (b *BlogImpl) List(ctx context.Context, filter models.Filter) ([]*models.Blog, error) {

	log.Printf("filter %+v", filter)
	filterData := bson.M{}
	if filter.AuthorID != "" {
		filterData["author_id"] = filter.AuthorID
	}
	if filter.Title != "" {
		filterData["title"] = bson.M{"$regex": filter.Title, "$options": "i"}
	}
	if !filter.CreatedAtGTE.IsZero() && !filter.CreatedAtGTE.Equal(time.Unix(0, 0)) {
		filterData["created_at"] = bson.M{"$gte": filter.CreatedAtGTE}
	}

	opts := options.Find()
	opts.SetSkip(int64(filter.Offset))
	opts.SetLimit(int64(filter.Limit))
	opts.SetCollation(&options.Collation{
		Locale:   "en",
		Strength: 2,
	})
	opts.SetSort(bson.M{"created_at": -1})

	cur, err := b.collection.Find(ctx, filterData, opts)
	if err != nil {
		return nil, err
	}

	blogs := make([]*models.Blog, 0, cur.RemainingBatchLength())

	for cur.Next(ctx) {
		var blogItem models.Blog
		err := cur.Decode(&blogItem)
		if err != nil {
			return nil, err
		}

		blogs = append(blogs, &blogItem)
	}
	// close the cursor
	if err := cur.Close(ctx); err != nil {
		return nil, err
	}

	return blogs, nil
}

// Update implements [Blog].
func (b *BlogImpl) Update(ctx context.Context, req *models.Blog) (*models.Blog, error) {

	data := bson.M{}
	if req.Title != "" {
		data["title"] = req.Title
	}
	if req.Content != "" {
		data["content"] = req.Content
	}
	if len(data) == 0 {
		return nil, ErrNoUpdatableFields
	}
	data["updated_at"] = time.Now()
	newBlog := &models.Blog{}
	err := b.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": req.ID},
		bson.M{"$set": data},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(newBlog)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrBlogNotFound
		}
		return nil, err
	}

	return newBlog, nil
}

// Delete implements [Blog].
func (b *BlogImpl) Delete(ctx context.Context, id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil || objId.IsZero() {
		return ErrInvalidID
	}

	res, err := b.collection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return ErrBlogNotFound
	}
	return nil
}
