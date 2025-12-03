package main

import (
	"fmt"

	"github.com/test/gorm_learn/code/constant"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`  //结构体标签
	Name  string `gorm:"column:name"` //gorm 标签
	Email string `gorm:"column:email"`
}

type Post struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"column:title"`
	Body   string `gorm:"column:body"`
	UserID uint   `gorm:"column:user_id"`
}

// InitDB 初始化数据库
func InitDB() *gorm.DB {
	db := ConnectDB()
	err := db.AutoMigrate(&User{}, &Post{})
	if err != nil {
		panic(err)
	}
	return db
}

// ConnectDB 连接数据库
func ConnectDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(constant.MYSQLDB), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	db := InitDB()
	fmt.Printf("jdbc 连接：%s\n", constant.MYSQLDB)

	// 插入数据
	db.Create(&User{Name: "Alice", Email: "alice@example.com"})
	db.Create(&User{Name: "Bob", Email: "bob@example.com"})
	db.Create(&User{Name: "Tom", Email: "tom@example.com"})

	//查询数据
	var users []User
	db.Find(&users)
	fmt.Println("所有用户: ", users)

	//更新数据
	db.Model(&users[0]).Update("Email", "alice@newexample.com")

	//删除数据
	db.Delete(&users[1])

	//查询剩余
	var remaining []User
	db.Find(&remaining)
	fmt.Println("剩余用户：", remaining)

	sqlDB, err := db.DB()
	if err != nil {
		panic("获取底层数据库连接失败：" + err.Error())
	}

	err = sqlDB.Close()
	if err != nil {
		panic(err)
	}

}
