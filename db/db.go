package db

import (
	"github.com/jinzhu/gorm"
)

type DbMyGoChat struct {
}

var dbMap = map[string]*gorm.DB{}

func GetDb(dbName string) (db *gorm.DB) {
	if db, ok := dbMap[dbName]; ok {
		return db
	} else {
		return nil
	}
}
