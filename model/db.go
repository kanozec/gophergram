package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitializeDB() (db *gorm.DB, err error) {
	db, err = gorm.Open("sqlite3", "./../..gophergram.db")
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxIdleConns(10)
	return db, nil
}
