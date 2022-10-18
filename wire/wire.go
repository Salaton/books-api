//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/Salaton/books-api/server/handlers"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func Wire(db *gorm.DB) *handlers.HandlersImpl {
	panic(wire.Build(ProviderSet))
}
