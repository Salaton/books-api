package database

import (
	"context"
	"fmt"

	"github.com/Salaton/books-api/models"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type BooksDB struct {
	db *gorm.DB
}

func NewBooksDataStore(DB *gorm.DB) *BooksDB {
	return &BooksDB{
		db: DB,
	}
}

func ConnectToDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=sala password=krychowiak-254 dbname=books port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}
	// export POSTGRESQL_URL="postgres://sala:krychowiak-254@localhost:5432/books?sslmode=disable"
	return db, nil
}

func (db *BooksDB) CreateComment(ctx context.Context, input models.Comments) error {
	comment := models.Comments{
		// TODO: create the UUID in a hook
		ID:        uuid.New().String(),
		Book:      input.Book,
		Comment:   input.Comment,
		IPAddress: input.IPAddress,
		CreatedAt: input.CreatedAt,
	}

	if err := db.db.Create(&comment).Error; err != nil {
		return fmt.Errorf("failed to create comment: %w", err)
	}

	return nil
}
