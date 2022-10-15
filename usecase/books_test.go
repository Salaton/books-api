package usecase_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Salaton/books-api/infrastructure/database/mock"
	"github.com/Salaton/books-api/models"
	"github.com/Salaton/books-api/usecase"
)

type BookStoreMock struct {
	GetBookDetailsFn   func(ctx context.Context) ([]*models.BookDetails, error)
	AddCommentFn       func(ctx context.Context, input models.Comments) error
	ListBookCommentsFn func(ctx context.Context, bookID string) ([]*models.Comments, error)
}

func (b *BookStoreMock) GetBookDetails(ctx context.Context) ([]*models.BookDetails, error) {
	return b.GetBookDetailsFn(ctx)
}

func (b *BookStoreMock) AddComment(ctx context.Context, input models.Comments) error {
	return b.AddCommentFn(ctx, input)
}

func (b *BookStoreMock) ListBookComments(ctx context.Context, bookID string) ([]*models.Comments, error) {
	return b.ListBookCommentsFn(ctx, bookID)
}

func TestBookStoreDetails_GetBookDetails(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []*models.BookDetails
		wantErr bool
	}{
		{
			name: "Happy Case - Successfully Get book details",
			args: args{
				ctx: context.Background(),
			},
			want: []*models.BookDetails{
				{
					URL:           "",
					Name:          "Just A Book",
					Isbn:          "123456",
					Authors:       []string{"Sala"},
					NumberOfPages: 260,
					Publisher:     "Ton Publishers",
					Country:       "Kenya",
					MediaType:     "PDF",
					Released:      "Today",
					Characters:    []string{"Sala", "Joe"},
					PovCharacters: []string{"Sala", "Joe"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := mock.PostgresMock{}
			mockBooks := BookStoreMock{}
			b := usecase.NewBookStoreImplementation(test, test)

			if tt.name == "Happy Case - Successfully Get book details" {
				mockBooks.GetBookDetailsFn = func(ctx context.Context) ([]*models.BookDetails, error) {
					return []*models.BookDetails{{
						URL:           "",
						Name:          "Just A Book",
						Isbn:          "123456",
						Authors:       []string{"Sala"},
						NumberOfPages: 260,
						Publisher:     "Ton Publishers",
						Country:       "Kenya",
						MediaType:     "PDF",
						Released:      "Today",
						Characters:    []string{"Sala", "Joe"},
						PovCharacters: []string{"Sala", "Joe"},
					}}, nil
				}
			}

			got, err := b.GetBookDetails(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("BookStoreDetails.GetBookDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("expected a response but got: %v", got)
				return
			}
		})
	}
}

func TestBookStoreDetails_AddComment(t *testing.T) {
	type args struct {
		ctx   context.Context
		input models.Comments
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy Case - Successfully add comments",
			args: args{
				ctx: context.Background(),
				input: models.Comments{
					ID:        "12",
					Book:      "Pragmatic Programmer",
					Comment:   "I really love this book",
					IPAddress: "123.23.23.23",
					CreatedAt: &time.Time{},
				},
			},
			wantErr: false,
		},
		{
			name: "Sad Case - Fail to add comments",
			args: args{
				ctx: context.Background(),
				input: models.Comments{
					ID:        "12",
					Book:      "Pragmatic Programmer",
					Comment:   "I really love this book",
					IPAddress: "123.23.23.23",
					CreatedAt: &time.Time{},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbMock := mock.NewPostgresMock()
			b := usecase.NewBookStoreImplementation(dbMock, dbMock)

			if tt.name == "Happy Case - Successfully add comments" {
				dbMock.CreateCommentFn = func(ctx context.Context, input models.Comments) error {
					return nil
				}
			}

			if tt.name == "Sad Case - Fail to add comments" {
				dbMock.CreateCommentFn = func(ctx context.Context, input models.Comments) error {
					return fmt.Errorf("failed to create comment")
				}
			}

			if err := b.AddComment(tt.args.ctx, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("BookStoreDetails.AddComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBookStoreDetails_ListBookComments(t *testing.T) {
	type args struct {
		ctx    context.Context
		bookID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy Case - Successfully list a book's comments",
			args: args{
				ctx:    context.Background(),
				bookID: "1",
			},
			wantErr: false,
		},
		{
			name: "Sad Case - Fail to list a book's comments",
			args: args{
				ctx:    context.Background(),
				bookID: "1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbMock := mock.NewPostgresMock()
			b := usecase.NewBookStoreImplementation(dbMock, dbMock)

			if tt.name == "Happy Case - Successfully list a book's comments" {
				dbMock.ListBookCommentsFn = func(ctx context.Context, bookID string) ([]*models.Comments, error) {
					return []*models.Comments{{
						ID:        "1",
						Book:      "That Book",
						Comment:   "Very lovely",
						IPAddress: "",
						CreatedAt: &time.Time{},
					}}, nil
				}
			}

			if tt.name == "Sad Case - Fail to list a book's comments" {
				dbMock.ListBookCommentsFn = func(ctx context.Context, bookID string) ([]*models.Comments, error) {
					return nil, fmt.Errorf("failed to list the book's comments")
				}
			}

			got, err := b.ListBookComments(tt.args.ctx, tt.args.bookID)
			if (err != nil) != tt.wantErr {
				t.Errorf("BookStoreDetails.ListBookComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("expected a response but got: %v", got)
				return
			}
		})
	}
}
