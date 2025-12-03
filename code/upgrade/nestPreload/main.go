package main

import (
	"fmt"

	"github.com/test/gorm_learn/code/constant"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/***
多级 Preload
在刚才 User -> Orders 的基础上，
我们再在 Order 下面加一层，比如叫 OrderItem：

一个订单 (Order) 有多个商品项 (OrderItem)

每个商品项有 Name 和 Quantity

这样就有了 User -> Orders -> OrderItems 的结构。
*/

type User struct {
	ID     uint `gorm:"primaryKey"`
	Name   string
	Email  string
	Orders []Order
}

type Order struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint // belongs to User
	Item       string
	Price      int
	OrderItems []OrderItem // 新增！一对多
}

type OrderItem struct {
	ID       uint   `gorm:"primaryKey"`
	OrderID  uint   // belongs to Order
	Name     string // 商品项名称
	Quantity int    // 数量
}

func insertTestData(db *gorm.DB) {
	var count int64
	db.Model(&OrderItem{}).Count(&count)
	if count > 0 {
		return
	}

	users := []User{
		{
			Name:  "Alice",
			Email: "alice@example.com",
			Orders: []Order{
				{
					Item:  "iPhone",
					Price: 999,
					OrderItems: []OrderItem{
						{Name: "iPhone 主机", Quantity: 1},
						{Name: "iPhone 保护壳", Quantity: 1},
					},
				},
				{
					Item:  "MacBook",
					Price: 1999,
					OrderItems: []OrderItem{
						{Name: "MacBook 主机", Quantity: 1},
					},
				},
			},
		},
		{
			Name:  "Bob",
			Email: "bob@example.com",
			Orders: []Order{
				{
					Item:  "AirPods",
					Price: 199,
					OrderItems: []OrderItem{
						{Name: "AirPods 耳机", Quantity: 1},
					},
				},
			},
		},
	}

	for _, user := range users {
		db.Create(&user)
	}
}

func main() {

	db, err := gorm.Open(mysql.Open(constant.MYSQLDB), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败：" + err.Error())
	}

	// 自动迁移：确保表存在
	err = db.AutoMigrate(&User{}, &Order{}, &OrderItem{})
	if err != nil {
		panic("迁移数据库失败：" + err.Error())
	}

	// 🔥 插入一些测试数据
	insertTestData(db)

	var users []User
	err = db.
		Preload("Orders").
		Preload("Orders.OrderItems").
		Find(&users).Error
	if err != nil {
		panic("查询失败：" + err.Error())
	}

	for _, user := range users {
		fmt.Printf("用户: %s (%s)\n", user.Name, user.Email)
		for _, order := range user.Orders {
			fmt.Printf("  - 订单: %s，价格: %d\n", order.Item, order.Price)
			for _, item := range order.OrderItems {
				fmt.Printf("      * 商品项: %s，数量: %d\n", item.Name, item.Quantity)
			}
		}
	}

}

/***

小总结 🧠

你想要的数据层级	写法
User -> Orders	Preload("Orders")
User -> Orders -> OrderItems	Preload("Orders.OrderItems")
User -> Orders -> OrderItems -> XXX	Preload("Orders.OrderItems.XXX")
可以一直多级 preload 下去，
只要模型关系定义正确（外键匹配）就行。
*/
