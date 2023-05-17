// this deals with connection with the sql database
package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// create a variable db
// the main intention of this file is to return a var db which other files will
// interact with

var (
	db *gorm.DB
)

// create a connect method which will help us to connect with the mysql db:

func Connect() {
	// open a database
	// gorm.Open("db_used", "username:pw@/dbname?")
	d, err := gorm.Open("mysql", "dipankar:dru563@/simplerest?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	db = d
}

// this func returns a db
func GetDB() *gorm.DB {
	return db
}
