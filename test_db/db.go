package test_db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewTestDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", "test_db", "user", "password", "book_alloc_test", "5432")
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	tx := d.Begin()
	tx.SavePoint("initial")
	defer tx.RollbackTo("initial")
	return tx, nil
}
