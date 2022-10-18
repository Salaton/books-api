// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/Salaton/books-api/infrastructure/database"
	"github.com/Salaton/books-api/server/handlers"
	"github.com/Salaton/books-api/usecase"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func Wire(db *gorm.DB) *handlers.HandlersImpl {
	booksDB := database.NewBooksDataStore(db)
	bookStoreDetails := usecase.NewBookStoreImplementation(booksDB, booksDB)
	handlersImpl := handlers.NewHandlersImplementation(bookStoreDetails)
	return handlersImpl
}
