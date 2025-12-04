package lesson02

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Name = u.Name + "_12123"
	return
}

type CreditCard struct {
	gorm.Model
	Number     string
	BankUserID uint
}

type BankUser struct {
	gorm.Model
	Name       string
	CreditCard CreditCard
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&BankUser{})
	db.AutoMigrate(&CreditCard{})

	// user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	//
	// result := db.Create(&user) // 通过数据的指针来创建
	// fmt.Println(result.RowsAffected)
	// user.ID             // 返回插入数据的主键
	// result.Error        // 返回 error
	// result.RowsAffected // 返回插入记录的条数

	// users := []*User{
	// 	{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
	// 	{Name: "Jackson", Age: 19, Birthday: time.Now()},
	// }

	// db.Create(users)
	// fmt.Println(users[0])

	// db.Create(&BankUser{
	// 	Name:       "jinzhu",
	// 	CreditCard: CreditCard{Number: "411111111111"},
	// })

	// user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	// db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)
	// user := User{Name: "Jinzhu", Age: 19, Birthday: time.Now()}
	// user.ID = 1
	// db.Clauses(clause.OnConflict{
	// 	Columns:   []clause.Column{{Name: "name"}},
	// 	DoUpdates: clause.AssignmentColumns([]string{"age"}),
	// }).Create(&user)

	// var user User
	// db.Debug().First(&user)
	// err := db.Debug().Take(&user).Error
	// if err != nil {
	// 	panic(err)
	// }

	// user.ID = 1
	// db.Debug().First(&user, []int{1, 2, 3})

	// var users []User
	// db.Debug().Find(&users)

	var user User
	user.ID = 1
	db.Debug().First(&user, "name = ?", "123123")
}
