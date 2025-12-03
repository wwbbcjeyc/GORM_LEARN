package main

import (
	"fmt"

	"github.com/test/gorm_learn/code/constant"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
解释一下这段：

概念	说明
User struct 里有 Orders []Order	GORM 根据字段名 Orders 自动建立外键关联
Preload("Orders")	查询 User 的同时把每个用户的 Orders 查出来
insertTestData 函数	保证只插一次测试数据，避免每次都重复插入
关闭数据库	db.DB().Close()，防止数据没写进去
*/

type User struct {
	ID     uint `gorm:"primaryKey"`
	Name   string
	Email  string
	Orders []Order // 关联字段，一对多
}

type Order struct {
	ID     uint   `gorm:"primaryKey"`
	UserID uint   // 外键字段
	Item   string // 商品名称
	Price  int    // 商品价格
}

func main() {
	db, err := gorm.Open(mysql.Open(constant.MYSQLDB), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败：" + err.Error())
	}

	// 自动迁移：确保表存在
	err = db.AutoMigrate(&User{}, &Order{})
	if err != nil {
		panic("迁移数据库失败：" + err.Error())
	}

	// 🔥 插入一些测试数据
	insertTestData(db)

	// 🌟 使用 Preload 查询用户及其订单
	var users []User
	err = db.Preload("Orders").Find(&users).Error
	if err != nil {
		panic("查询失败：" + err.Error())
	}

	// 打印查询结果
	for _, user := range users {
		fmt.Printf("用户: %s (%s)\n", user.Name, user.Email)
		for _, order := range user.Orders {
			fmt.Printf("  - 订单: %s，价格: %d\n", order.Item, order.Price)
		}
	}

	// 关闭连接
	sqlDB, err := db.DB()
	if err != nil {
		panic("获取底层数据库连接失败：" + err.Error())
	}
	sqlDB.Close()
}

func insertTestData(db *gorm.DB) {
	// 检查是否已有数据
	var count int64
	db.Model(&Order{}).Count(&count)
	if count > 0 {
		return // 如果已有数据，不重复插入
	}

	// 创建订单数据
	users := []User{
		{
			Name:  "Alice",
			Email: "alice@example.com",
			Orders: []Order{
				{Item: "iPhone", Price: 999},
				{Item: "MacBook", Price: 1999},
			},
		},
		{
			Name:  "Bob",
			Email: "bob@example.com",
			Orders: []Order{
				{Item: "AirPods", Price: 199},
			},
		},
	}

	// 批量插入
	for _, user := range users {
		db.Create(&user)
	}
}
