package db

import (
	"database/sql/driver"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Context struct {
	db *gorm.DB
}

func NewDB() (*gorm.DB, error) {
	dsn := "host=db user=user password=password dbname=book_alloc port=5432 sslmode=disable TimeZone=Asia/Tokyo"
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
