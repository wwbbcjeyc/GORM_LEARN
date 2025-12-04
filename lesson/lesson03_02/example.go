package lesson0302

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	CreditCards []CreditCard
}

// type User struct {
// 	gorm.Model
// 	Name       string     `gorm:"unique;index;size:50"`
// 	CreditCard CreditCard `gorm:"foreignKey:UserName;references:Name"`
// }

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&CreditCard{})

	// user := User{}
	// db.Create(&user)

	// card := CreditCard{Number: "1111", UserID: 1}
	// db.Create(&card)
	//
	// card2 := CreditCard{Number: "2222", UserID: 1}
	// db.Create(&card2)

	user := User{}
	db.Preload("CreditCards").First(&user, 1)
	fmt.Println(user)
}
