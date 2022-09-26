package repository

import (
	"context"

	"github.com/Salaton/books-api/models"
)

type Create interface {
	CreateComment(ctx context.Context, input models.Comments) error
}
