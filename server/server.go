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

	postgresDB, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatal().Err(err).Msg("an error occured while connecting to the database")
	}

	repository := database.NewBooksDataStore(postgresDB)
	booksUseCase := usecase.NewBookStoreImplementation(repository)
	handler := handlers.NewHandlersImplementation(booksUseCase)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/books", handler.GetBookDetails)
		v1.POST("/comments", handler.AddComment)
	}

	return router
}
