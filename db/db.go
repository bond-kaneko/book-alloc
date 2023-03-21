package db

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

type Context struct {
	db *gorm.DB
}

func NewDB() (*gorm.DB, error) {
	e := os.Getenv("ENV")
	if e == "" || e == "test" {
		panic("Use NewTestDB() in the test environment")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func NewTestDB() (*gorm.DB, error) {
	e := os.Getenv("ENV")
	if e != "test" {
		panic("Use NewDB() except in the test environment")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", "test_db", "user", "password", "book_alloc_test", "5432")
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

type NullableTime struct {
	Time  time.Time
	Valid bool
}

func (nt *NullableTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

func (nt *NullableTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}
