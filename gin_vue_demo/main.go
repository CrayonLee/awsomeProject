package main

import (
	"gin_vue/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	InitConfig()
	 common.InitDB()


	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}

func InitConfig()  {
	//获取当前工作路径
	wordDir, err := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(wordDir+"/config")
	err = viper.ReadInConfig()
	if err != nil {
		panic("配置文件读取失败"+err.Error())

	}

}
