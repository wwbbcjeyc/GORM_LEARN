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

	// 调用不同的函数
	batchQuery(db)
	batchUpdate(db)
	batchDelete(db)

}

// 批量查询
func batchQuery(db *gorm.DB) {
	var users []User
	err := db.Where("email LIKE ?", "%example.com").Find(&users).Error
	if err != nil {
		panic("批量查询失败：" + err.Error())
	}

	for _, user := range users {
		fmt.Printf("用户: ID=%d, 名字=%s, 邮箱=%s\n", user.ID, user.Name, user.Email)
	}

}

// 批量更新
func batchUpdate(db *gorm.DB) {
	err := db.Model(&User{}).Where("name = ?", "Alice").Update("email", "alice@newdomain.com").Error

	if err != nil {
		panic("批量更新失败：" + err.Error())
	}
	fmt.Println("【批量更新】完成")
}

// 批量删除
func batchDelete(db *gorm.DB) {
	// 删除价格低于 500 的订单
	err := db.Where("price < ?", 500).Delete(&Order{}).Error
	if err != nil {
		panic("批量删除失败：" + err.Error())
	}

	fmt.Println("【批量删除】完成")

}
