package main

import (
	"fmt"
	"log"

	"github.com/test/gorm_learn/code/constant"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(mysql.Open(constant.MYSQLDB), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 根据结构体自动创建或更新数据库表结构
	db.AutoMigrate(&Product{})

	// 创建记录

	result := db.Create(&Product{Code: "D42", Price: 100})
	if result.Error != nil {
		log.Fatalf("创建失败: %v", result.Error)
	}

	// 查询记录
	//var product Product
	//db.First(&product, 1)                 // // 查询 ID=1 的记录
	//db.First(&product, "code = ?", "D42") //查询 code='D42' 的记录

	// 获取创建的记录ID
	var product Product
	productID := product.ID
	fmt.Printf("新创建记录的ID: %d\n", productID)

	fmt.Println("=== 查询刚创建的记录 ===")
	db.First(&product, productID) // 使用变量ID
	fmt.Printf("查询结果1: ID=%d, Code=%s, Price=%d\n", product.ID, product.Code, product.Price)
	var product2 Product
	db.First(&product2, "code = ?", "D42")
	fmt.Printf("查询结果2: ID=%d, Code=%s, Price=%d\n", product2.ID, product2.Code, product2.Price)

	// 更新记录 更新单个字段
	//db.Model(&product).Update("Price", 200)
	// 更新多个字段（只更新非零值字段）
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	// 更新多个字段（使用map，更新零值字段）
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Update - 更新记录
	fmt.Println("=== 更新记录 ===")
	db.Model(&product).Update("Price", 200)

	// 验证更新结果
	db.First(&product, productID)
	fmt.Printf("更新后: ID=%d, Code=%s, Price=%d\n", product.ID, product.Code, product.Price)

	// 删除记录（Delete）
	//db.Delete(&product, 1)

	// Delete - 删除记录（软删除）
	fmt.Println("=== 删除记录 ===")
	db.Delete(&product, productID)

	// 查询所有记录（包括未删除的）
	var products []Product
	db.Find(&products)
	fmt.Printf("当前表中记录数: %d\n", len(products))
}
