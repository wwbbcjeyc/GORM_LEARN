package lesson03

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company
}

type UserB struct {
	gorm.Model
	Name       string
	CompanyIDS int
	Company    Company `gorm:"foreignKey:CompanyIDS"`
}

type UserC struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company `gorm:"references:CardNumber"`
}

type Company struct {
	ID         int
	CardNumber int `gorm:"unique"`
	Name       string
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{})
	// db.AutoMigrate(&UserB{})
	// db.AutoMigrate(&UserC{})
	db.AutoMigrate(&Company{})

	// user := User{Name: "user normal", CompanyID: 1}
	// db.Create(&user)

	// company := Company{Name: "John com", CardNumber: 1111}
	// db.Create(&company)

	// user := User{}
	// db.First(&user)
	// fmt.Println(user)

	// user := UserB{CompanyIDS: 1, Name: "user_reference"}
	// db.Create(&user)

	user := User{}
	db.Preload("Company").First(&user)
	fmt.Println(user)

	// user := UserC{CompanyID: 1111, Name: "user_cccc"}
	// db.Create(&user)
}
