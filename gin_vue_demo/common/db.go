package common

import (
	"fmt"
	"gin_vue/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
)
var DB *gorm.DB
func InitDB() *gorm.DB {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))
	fmt.Printf(args)
	db, err := gorm.Open(mysql.Open(args),&gorm.Config{})
	if err != nil {
		panic("fail to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&model.User{})

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}