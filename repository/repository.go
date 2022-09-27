package repository

import (
	"context"

	"github.com/Salaton/books-api/models"
)

type Create interface {
	CreateComment(ctx context.Context, input models.Comments) error
}

type Query interface {
	ListBookComments(ctx context.Context, bookID string) ([]*models.Comments, error)
}
