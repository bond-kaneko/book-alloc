package db

import (
	"gorm.io/gorm"
)

type Context struct {
	db *gorm.DB
}

func NewDB() (*Context, error) {
	// TODO 実体を返す
	return nil, nil
}

// TODO Selectとかのメソッドを実装
