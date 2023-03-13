package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Context struct {
	db *gorm.DB
}

func NewDB() (*gorm.DB, error) {
	dsn := "host=db user=user password=password dbname=book_allock port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
