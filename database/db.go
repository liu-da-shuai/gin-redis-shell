package database

import (
	"gorm.io/gorm"
	"sync"
)

var (
	once sync.Once
	DB   *db
)

type db struct {
	Db *gorm.DB
}

func newDb() *db {
	once.Do(func() {
		DB = &db{}
	})
	return DB // todo 实现
}
