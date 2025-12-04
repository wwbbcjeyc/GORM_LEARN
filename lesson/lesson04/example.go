package lesson04

import (
	"fmt"

	"gorm.io/gorm"
)

type Dog struct {
	ID   int
	Name string
	// Toy  Toy `gorm:"polymorphic:Owner"`
	Toy Toy `gorm:"polymorphicType:Kind;polymorphicId:TID;polymorphicValue:dog"`
}

type Cat struct {
	ID   int
	Name string
	Age  int
	// Toy  Toy `gorm:"polymorphic:Owner"`
	Toy Toy `gorm:"polymorphicType:Kind;polymorphicId:TID;polymorphicValue:cat"`
}

type Toy struct {
	ID   int
	Name string
	Kind string
	TID  int
}

func (c *Cat) AfterDelete(tx *gorm.DB) error {
	fmt.Println(c)
	return nil
}

func (c *Cat) BeforeDelete(tx *gorm.DB) error {
	fmt.Println(c)
	return nil
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&Dog{}, &Cat{}, &Toy{})

	// 多态
	db.Create(&Dog{Name: "wangcai", Toy: Toy{Name: "gutou"}})
	db.Create(&Cat{Name: "mimi", Toy: Toy{Name: "doumaobang"}, Age: 1})
	db.Create(&Cat{Name: "mimi2", Toy: Toy{Name: "doumaobang"}, Age: 2})
	db.Create(&Cat{Name: "mimi3", Toy: Toy{Name: "doumaobang"}, Age: 3})

	// var dog Dog
	// var cat Cat
	// db.Preload("Toy").First(&dog)
	// db.Preload("Toy").First(&cat)
	// fmt.Println(dog, cat)

	// RunHook(db)
	// RunTransaction(db)
	// RunDefinition(db)

	db.Debug().Where("age > ?", 2).Delete(&Cat{})
}
