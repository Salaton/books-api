package server

import (
	"context"

	"github.com/Salaton/books-api/server/handlers"
	"github.com/Salaton/books-api/usecase"
	"github.com/gin-gonic/gin"
)

func Router(ctx context.Context) *gin.Engine {
	router := gin.Default()

	booksUseCase := usecase.NewBookStoreImplementation()
	handler := handlers.NewHandlersImplementation(booksUseCase)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/books", handler.GetBookDetails)
	}

	return router
}
