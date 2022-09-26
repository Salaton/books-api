package server

import (
	"context"

	"github.com/Salaton/books-api/infrastructure/database"
	"github.com/Salaton/books-api/server/handlers"
	"github.com/Salaton/books-api/usecase"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Router(ctx context.Context) *gin.Engine {
	router := gin.Default()

	booksUseCase := usecase.NewBookStoreImplementation()
	handler := handlers.NewHandlersImplementation(booksUseCase)

	_, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatal().Err(err).Msg("an error occured while connecting to the database")
	}

	v1 := router.Group("/api/v1")
	{
		v1.GET("/books", handler.GetBookDetails)
	}

	return router
}
