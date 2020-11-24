package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	err := initDB()
	if err!=nil{
		panic(err)
	}
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	
	r.GET("/book/list", bookListHandler)
	r.GET("/book/toAdd",toAddBookHandler)
	r.POST("/book/add",addBookHandler)
	r.GET("/book/delete/:id",deleteHandler)
	r.GET("/book/edit/:id",toEditHandler)
	r.POST("/book/edit",editHandler)
	r.Run()
}

func editHandler(context *gin.Context) {
	var msg string
	idStr := context.PostForm("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err!=nil{
		msg="数据格式转换失败"
		context.JSON(http.StatusOK,gin.H{
			"msg":msg,
		})
	}
	bookName := context.PostForm("book_name")
	priceStr := context.PostForm("price")
	price, err := strconv.ParseFloat(priceStr, 10)
	if err!=nil{
		msg="数据格式转换失败"
		context.JSON(http.StatusOK,gin.H{
			"msg":msg,
		})
	}
	err = modify("update book set book_name=?,price=? where id=?", bookName, price, id)
	if err!=nil{
		msg="更新失败"
		context.JSON(http.StatusOK,gin.H{
			"msg":msg,
		})
	}
	//更新成功  跳转至图书列表页面
	context.Redirect(http.StatusMovedPermanently,"/book/list")
}

func toEditHandler(context *gin.Context) {
	var msg string
	param := context.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err!=nil{
		msg="数据格式转换失败"
		context.JSON(http.StatusOK,gin.H{
			"msg":msg,
		})
	}
	one, err := queryOne("select id,book_name,price from book where id=?", id)
	if err!=nil{
		fmt.Printf("err:%v\n",err)
		msg="根据id查询失败"
		context.JSON(http.StatusOK,gin.H{
			"msg":msg,
		})
	}
	context.HTML(http.StatusOK,"book/edit.html",gin.H{
		"code":1,
		"msg":msg,
		"data":one,
	})
}

func deleteHandler(context *gin.Context) {
	var msg string
	param := context.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err!=nil{
		msg="您输入的数字不合法"
		context.JSON(http.StatusOK,gin.H{
			"msg":msg,
		})
	}
	err = modify("delete from book where id=?", id)
	if err!=nil{
		msg="删除失败"
		context.JSON(http.StatusOK,gin.H{
			"msg":msg,
		})
	}
	//删除成功  重定向至图书列表界面
	context.Redirect(http.StatusMovedPermanently,"/book/list")

}

func addBookHandler(context *gin.Context) {
	var msg string
	bookName := context.PostForm("book_name")
	priceStr := context.PostForm("price")
	float, err := strconv.ParseFloat(priceStr, 64)
	if err!=nil{
		msg="无效的价格参数"
		context.JSON(http.StatusOK,gin.H{
			"msg":msg,
		})
	}
	err = modify("insert into book(book_name,price) values (?,?)", bookName, float)
	if err!=nil{
		msg="添加失败"
		context.JSON(http.StatusOK,gin.H{
			"msg":msg,
		})
	}
	//重定向值图书列表页面
	context.Redirect(http.StatusMovedPermanently,"/book/list")

}

func toAddBookHandler(context *gin.Context) {
	context.HTML(http.StatusOK,"book/new_book.html",nil)
}

func bookListHandler(context *gin.Context) {
	list, err := queryMulti("select id,book_name,price from book")
	if err!=nil{
		context.JSON(http.StatusOK,gin.H{
			"code":0,
			"msg":"err",
		})
	}
	context.HTML(http.StatusOK,"book/list.html",gin.H{
		"code":1,
		"msg":"查询成功",
		"data":list,
	})
}
