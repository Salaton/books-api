package wire

import (
	"github.com/Salaton/books-api/infrastructure/database"
	"github.com/Salaton/books-api/repository"
	"github.com/Salaton/books-api/server/handlers"
	"github.com/Salaton/books-api/usecase"
	"github.com/google/wire"
)

// Provider is an initializer function where you create a single struct. For example:
// func NewBooksDataStore(DB *gorm.DB) *BooksDB {
// 	return &BooksDB{
// 		db: DB,
// 	}
// }

// This is used in context of the package `Wire` which is used for Dependency Injection

// ProviderSet is a set of providers grouped into one
var ProviderSet wire.ProviderSet = wire.NewSet(
	usecase.NewBookStoreImplementation,
	handlers.NewHandlersImplementation,
	database.NewBooksDataStore,

	wire.Bind(new(usecase.BookStore), new(*usecase.BookStoreDetails)),
	wire.Bind(new(handlers.Handlers), new(*handlers.HandlersImpl)),
	wire.Bind(new(repository.Create), new(*database.BooksDB)),
	wire.Bind(new(repository.Query), new(*database.BooksDB)),
)
