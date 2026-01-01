package repo

import (
	"context"

	"github.com/MidhunJithu/grpc-go-sample/blog/server/models"
)

type Blog interface {
	Create(context.Context, *models.Blog) error
	Read(context.Context, string) (*models.Blog, error)
}
