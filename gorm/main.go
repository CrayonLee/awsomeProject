package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)



type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivedAt    sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}



func main() {
	db, err := gorm.Open("mysql", "root:lzh123456@(39.102.112.15:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err!=nil{
		panic(err)
	}
	defer db.Close()

	/*// 自动迁移
	db.SingularTable(true)
	db.AutoMigrate(&User{})

	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	 //db.Create(&user) // 通过数据的指针来创建
	db.Select("Name", "Age", "CreatedAt").Create(&user)*/


	//db.Model(&User{}).Create(map[string]interface{}{
	//	"Name": "jinzhu", "Age": 18,
	//})

	db.Model(&User{}).Create([]map[string]interface{}{
		{"Name":"test1","Age":18},
		{"Name":"test2","Age":18},
	})

}