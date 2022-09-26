package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Salaton/books-api/models"
)

const (
	// TODO: Should be in a config
	booksAPIURL = "https://anapioficeandfire.com/api/books"
)

type BookStore interface {
	GetBookDetails(ctx context.Context) ([]*models.BookDetails, error)
}

type BookStoreDetails struct {
}

func NewBookStoreImplementation() *BookStoreDetails {
	return &BookStoreDetails{}
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
