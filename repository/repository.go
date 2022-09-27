package repository

import (
	"context"

	"github.com/Salaton/books-api/models"
)

type Create interface {
	CreateComment(ctx context.Context, input models.Comments) error
	ListBookComments(ctx context.Context, bookID string) ([]*models.Comments, error)
}
