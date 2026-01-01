package repo

import (
	"context"

	"github.com/MidhunJithu/grpc-go-sample/blog/server/models"
)

type Blog interface {
	Create(context.Context, *models.Blog) error
	Read(context.Context, string) (*models.Blog, error)
	List(context.Context, models.Filter) ([]*models.Blog, error)
	Update(context.Context, *models.Blog) (*models.Blog, error)
	Delete(context.Context, string) error
}

type RepoError struct {
	Code    int
	Message string
}

func (e *RepoError) Error() string {
	return e.Message
}

var (
	ErrBlogNotFound = &RepoError{
		Code:    404,
		Message: "blog not found",
	}
	ErrNoUpdatableFields = &RepoError{
		Code:    400,
		Message: "no updatable fields",
	}
	ErrInvalidID = &RepoError{
		Code:    400,
		Message: "invalid id",
	}
)
