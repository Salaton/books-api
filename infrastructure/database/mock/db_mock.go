package mock

import (
	"context"
	"time"

	"github.com/Salaton/books-api/models"
)

type PostgresMock struct {
	CreateCommentFn    func(ctx context.Context, input models.Comments) error
	ListBookCommentsFn func(ctx context.Context, bookID string) ([]*models.Comments, error)
}

func (p PostgresMock) CreateComment(ctx context.Context, input models.Comments) error {
	return p.CreateCommentFn(ctx, input)
}

func (p PostgresMock) ListBookComments(ctx context.Context, bookID string) ([]*models.Comments, error) {
	return p.ListBookCommentsFn(ctx, bookID)
}

func NewPostgresMock() *PostgresMock {
	return &PostgresMock{
		CreateCommentFn: func(ctx context.Context, input models.Comments) error {
			return nil
		},
		ListBookCommentsFn: func(ctx context.Context, bookID string) ([]*models.Comments, error) {
			return []*models.Comments{{
				ID:        bookID,
				Book:      bookID,
				Comment:   "",
				IPAddress: "",
				CreatedAt: &time.Time{},
			}}, nil
		},
	}
}
