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
	UserID uint // 外键
	Item   string
	Price  int
}

func main() {
	db, err := gorm.Open(mysql.Open(constant.MYSQLDB), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败：" + err.Error())
	}

	// 只演示 joins，不动原有表结构
	joinDemo(db)
}

func joinDemo(db *gorm.DB) {
	type Result struct {
		UserName   string
		UserEmail  string
		OrderItem  string
		OrderPrice int
	}

	var results []Result

	err := db.Table("users").
		Select("users.name as user_name, users.email as user_email, orders.item as order_item, orders.price as order_price").
		Joins("left join orders on orders.user_id = users.id").
		Scan(&results).Error
	if err != nil {
		panic("查询失败：" + err.Error())
	}

	// 打印结果
	for _, r := range results {
		fmt.Printf("用户: %s (%s)，购买: %s，价格: %d\n", r.UserName, r.UserEmail, r.OrderItem, r.OrderPrice)
	}
}
