package main

import (
	"github.com/test/gorm_learn/code/constant"
	"github.com/test/gorm_learn/lesson/lesson01"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Parent struct {
	ID   int `gorm:"primary_key"`
	Name string
}

type Child struct {
	Parent
	Age int
}

func InitDB(dst ...interface{}) *gorm.DB {
	db, err := gorm.Open(mysql.Open(constant.MYSQLDB))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(dst...)

	return db
}

func main() {
	db, err := gorm.Open(mysql.Open(constant.MYSQLDB))
	if err != nil {
		panic(err)
	}

	lesson01.Run(db)
	// lesson02.Run(db)
	// lesson03.Run(db)
	// lesson0302.Run(db)
	// lesson0303.Run(db)
	// lesson0304.Run(db)
	// lesson04.Run(db)

	InitDB(&Parent{}, &Child{})
}
