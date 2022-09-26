package database

import (
	"fmt"

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
	return db, nil
}
