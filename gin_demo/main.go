package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func test1(){
	//创建一个默认的路由引擎
	r := gin.Default()
	//Get：请求方式；/hello：请求路径
	//当客户端以Get方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(context *gin.Context) {
		//context.JSON 会返回JSON格式的数据
		context.JSON(http.StatusOK,gin.H{
			"message":"Hello World",
		})
	})
	//启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}

func htmlRender(){
	r:=gin.Default()
	//自定义模板函数
	r.LoadHTMLGlob("templates/**")
	//r.SetFuncMap(template.FuncMap{
	//	"safe": func(str string) template.HTML {
	//		return template.HTML(str)
	//	},
	//})
	//r.LoadHTMLFiles("templates/index.tmpl")
	//r.GET("/index", func(context *gin.Context) {
	//	context.HTML(http.StatusOK,"index.tmpl","<a href='http://www.liwenzhou.com'>李文周的博客</a>")
	//
	//})
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK,"index.html",gin.H{
			"title":"index---",
		})
	})
	r.GET("/msg", func(context *gin.Context) {
		context.HTML(http.StatusOK,"msg.html",gin.H{
			"title":"context---",
		})

	})



	r.Static("/static","./static")
	r.Run()
}

//获取query_string参数
func queryStirng()  {
	r:=gin.Default()
	r.GET("/user/search", func(context *gin.Context) {
		username:=context.DefaultQuery("username","小王子")
		address:=context.Query("address")
		context.JSON(http.StatusOK,gin.H{
			"message":"Ok",
			"username":username,
			"address":address,
		})
	})
	r.Run()
}

//获取表单参数
func queryForm()  {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.POST("/user/search", func(c *gin.Context) {
		// DefaultPostForm取不到值时会返回指定的默认值
		//username := c.DefaultPostForm("username", "小王子")
		username := c.PostForm("username")
		address := c.PostForm("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	r.Run(":8080")
}

//获取路径参数
func queryPath()  {
	r:=gin.Default()
	r.GET("/user/search/:username/:address", func(context *gin.Context) {
		username:=context.Param("username")
		password:=context.Param("password")
		//将json结果输出
		context.JSON(http.StatusOK,gin.H{
			"message":"ok",
			"username":username,
			"password":password,
		})
	})
	r.Run()
}
func main() {
	//queryStirng()
	//queryPath()

	//gin框架操作cookie
	r := gin.Default()

	r.GET("/cookie", func(context *gin.Context) {

		cookie, err := context.Cookie("gin_cookie")
		if err!=nil{
			cookie="NotSet"
			//设置cookie
			context.SetCookie("gin_cookie","test",3600,"/","localhost",false,true)
		}
		fmt.Printf("cookie value is %s \n",cookie)
	})
	r.Run()
}
