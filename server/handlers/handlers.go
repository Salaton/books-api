package handlers

import (
	"context"
	"net/http"

	"github.com/Salaton/books-api/usecase"
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	GetBookDetails(c *gin.Context)
}

type HandlersImpl struct {
	books usecase.BookStore
}

func NewHandlersImplementation(books usecase.BookStore) *HandlersImpl {
	return &HandlersImpl{
		books: books,
	}
}

func jsonErrorResponse(c *gin.Context, statusCode int, err string) {
	c.JSON(statusCode, gin.H{"error": err})
}

func (h HandlersImpl) GetBookDetails(c *gin.Context) {
	ctx := context.Background()
	books, err := h.books.GetBookDetails(ctx)
	if err != nil {
		jsonErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"books": books})
}
