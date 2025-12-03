package main

import (
	"fmt"

	"github.com/test/gorm_learn/code/constant"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
}

type Order struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	Item   string
	Price  int
}

func main() {
	db, err := gorm.Open(mysql.Open(constant.MYSQLDB), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败：" + err.Error())
	}

	transactionDemo(db)
}

func transactionDemo(db *gorm.DB) {
	err := db.Transaction(func(tx *gorm.DB) error {
		// 1. 批量更新
		if err := tx.Model(&User{}).
			Where("name = ?", "Bob").
			Update("email", "bob@newdomain.com").Error; err != nil {
			return err // 出错自动回滚
		}

		// 2. 批量删除
		if err := tx.Where("price < ?", 500).
			Delete(&Order{}).Error; err != nil {
			return err // 出错自动回滚
		}

		// 3. 手动插入一条新用户
		newUser := User{Name: "Charlie", Email: "charlie@example.com"}
		if err := tx.Create(&newUser).Error; err != nil {
			return err // 出错自动回滚
		}

		// 所有操作都成功了，事务提交
		return nil
	})

	if err != nil {
		fmt.Println("事务执行失败，已回滚：", err)
	} else {
		fmt.Println("事务执行成功，已提交！")
	}
}
