package lesson0303

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Refer     uint       `gorm:"unique"`
	Languages []Language `gorm:"many2many:user_lang;foreignkey:Refer;joinForeignKey:UserReferID;References:Refer;joinReferences:LangReferID;"`
}

// type User struct {
// 	gorm.Model
// 	Languages []Language
// }

type Language struct {
	gorm.Model
	Name  string
	Refer uint `gorm:"unique"`
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Language{})

	// user := User{}
	// db.Create(&user)

	// language := Language{Name: "english"}
	// db.Create(&language)
	//
	// language2 := Language{Name: "chinese"}
	// db.Create(&language2)

	// user := User{Languages: []Language{{Name: "en1"}, {Name: "en2"}}}
	// db.Create(&user)

	// user := User{}
	// err := db.Preload("Languages").Find(&user, 3).Error
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(user)
}
