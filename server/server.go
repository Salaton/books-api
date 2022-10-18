package server

import (
	"context"

	"github.com/Salaton/books-api/infrastructure/database"
	"github.com/Salaton/books-api/wire"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Router(ctx context.Context) *gin.Engine {
	router := gin.Default()

	postgresDB, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatal().Err(err).Msg("an error occurred while connecting to the database")
	}
	handler := wire.Wire(postgresDB)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/books", handler.GetBookDetails)
		v1.POST("/comments/:bookID", handler.AddComment)
		v1.GET("/comments/:bookID", handler.ListBookComments)
	}

	return router
}
