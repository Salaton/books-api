package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Salaton/books-api/models"
	"github.com/Salaton/books-api/repository"
)

const (
	// TODO: Should be in a config
	booksAPIURL = "https://anapioficeandfire.com/api/books"
)

type BookStore interface {
	GetBookDetails(ctx context.Context) ([]*models.BookDetails, error)
	AddComment(ctx context.Context, input models.Comments) error
	ListBookComments(ctx context.Context, bookID string) ([]*models.Comments, error)
}

type BookStoreDetails struct {
	create repository.Create
}

func NewBookStoreImplementation(create repository.Create) *BookStoreDetails {
	return &BookStoreDetails{
		create: create,
	}
}

func (b *BookStoreDetails) GetBookDetails(ctx context.Context) ([]*models.BookDetails, error) {
	var booksData []*models.BookDetails

	response, err := MakeRequest(ctx, http.MethodGet, booksAPIURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %w", err)
	}

	err = json.Unmarshal(data, &booksData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall data: %w", err)
	}

	return booksData, nil
}

func (b *BookStoreDetails) AddComment(ctx context.Context, input models.Comments) error {
	err := b.create.CreateComment(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (b *BookStoreDetails) ListBookComments(ctx context.Context, bookID string) ([]*models.Comments, error) {
	comments, err := b.create.ListBookComments(ctx, bookID)
	if err != nil {
		return nil, fmt.Errorf("failed to list the book's comments: %w", err)
	}
	return comments, nil
}

// TODO: Extract this to a different file
func MakeRequest(ctx context.Context, method string, path string, body interface{}) (*http.Response, error) {
	client := http.Client{}

	encoded, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	payload := bytes.NewBuffer(encoded)
	req, err := http.NewRequestWithContext(ctx, method, path, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("an error occured while sending a HTTP request: %w", err)
	}

	return response, nil
}
