package lesson0304

import (
	"gorm.io/gorm"
)

type Address struct {
	ID       uint `gorm:"primarykey"`
	Address1 string
	UserID   uint
}

type Email struct {
	ID     uint `gorm:"primarykey"`
	Email  string
	UserID uint
}

type Language struct {
	ID   uint `gorm:"primarykey"`
	Name string
}

type Company struct {
	ID   uint `gorm:"primarykey"`
	Name string
}

type User struct {
	ID             uint `gorm:"primarykey"`
	Name           string
	BillingAddress Address
	// ShippingAddress Address
	Emails    []Email
	Languages []Language `gorm:"many2many:user_languages;"`
	CompanyID uint
	Company   Company
}

func CreateUser(db *gorm.DB) {
	comp := Company{Name: "Comp"}
	db.Create(&comp)

	user := User{
		Name:           "jinzhu11111",
		BillingAddress: Address{Address1: "Billing Address - Address 11111"},
		// ShippingAddress: Address{Address1: "Shipping Address - Address 111111"},
		Emails: []Email{
			{Email: "11@example.com"},
			{Email: "22@example.com"},
		},
		Languages: []Language{
			{Name: "ZH"},
			{Name: "EN"},
		},
		CompanyID: comp.ID,
	}
	db.Create(&user)
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Language{})
	db.AutoMigrate(&Email{})
	db.AutoMigrate(&Address{})
	db.AutoMigrate(&Company{})

	// CreateUser(db)

	// var user User
	// db.First(&user)
	// db.Preload("BillingAddress").Preload("Emails").First(&user)
	// db.Preload(clause.Associations).First(&user)
	// fmt.Println(user)

	// var langs []Language
	// db.Model(&User{ID: 1}).Association("Languages").Find(&langs)
	// fmt.Println(langs)

	// var user User
	// db.Preload(clause.Associations).First(&user)
	// // user.BillingAddress = Address{Address1: "111"}
	// user.BillingAddress.Address1 = "444"
	// // db.Debug().Updates(&user)
	// db.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)

	// 如果user中有2个Address，就会发生异常，因为2个地址不能明确的判定出对应关系，这个时候应该在结构中明确指定地址ID

	// 一对多更新
	// var emails []Email
	// db.Model(&User{ID: 1}).Association("Emails").Find(&emails)
	// // emails[0].Email = "1111@example.com"
	// // db.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Model(&User{ID: 1}).Association("Emails").Replace(emails)
	// db.Debug().Model(&User{ID: 1}).Association("Emails").Replace(&Email{Email: "11@11.com"}, &Email{Email: "22@22.com"})

	// 多对多更新
	// var langZH, langEN Language
	// db.First(&langZH, "name = ?", "ZH")
	// db.First(&langEN, "name = ?", "EN")
	// // db.Debug().Model(&User{ID: 1}).Association("Languages").Replace(&langZH) // 必须是引用
	// // db.Debug().Model(&User{ID: 1}).Association("Languages").Delete(langZH)
	// // db.Debug().Model(&User{ID: 1}).Association("Languages").Append(&langEN) // 必须是引用
	// // db.Debug().Model(&User{ID: 1}).Association("Languages").Append(&Language{Name: "FR"})
	// db.Debug().Model(&User{ID: 1}).Association("Languages").Clear()

	// 删除关联
	db.Debug().Select("Emails", "Languages", "BillingAddress").Delete(&User{ID: 1})
}
