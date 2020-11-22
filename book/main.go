package main

import (
	"fmt"
	"os"
)

/**
	需求：
		使用函数实现一个简单的图书馆里系统
		每本图书有书名 作者 价格 上架信息
		用户可以从控制台添加书籍 修改书籍信息 打印所有书籍列表

	需求分析
		0.定义结构体
		1.打印菜单
		2.等待用户输入菜单选项
		3.书写添加书籍的函数
		4.书写修改书籍的函数
		5.展示书籍的函数
		6.退出os.Exit(0)
 */

type book struct{
	title string
	author string
	price float32
	publish bool
}

//定义一个创建新书时使用的构造函数
func newBook(title,author string,price float32,publish bool) *book  {
	return &book{
		title: title,
		author: author,
		price: price,
		publish: publish,
	}
}

//定义一个用于存放所有书籍信息的切片
var allBooks=make([]*book,0,200)

//用于获取用户输入
func userInput() *book  {
	var (
		title string
		author string
		price float32
		publish bool

	)
	//1.获取用户输入
	fmt.Println("请根据提示输入相关内容")
	fmt.Print("请输入书名：")
	fmt.Scanln(&title)
	fmt.Print("请输入作者：")
	fmt.Scanln(&author)
	fmt.Print("请输入价格：")
	fmt.Scanln(&price)
	fmt.Print("请输入是否上架【true|false】：")
	fmt.Scanln(&publish)

	//2.创建一本新书
	book:=newBook(title,author,price,publish)
	//fmt.Printf("%v\n",book)
	return book
}

func showMenu()  {
	fmt.Println("欢迎进入Book Management System！")
	fmt.Println("1.添加书籍")
	fmt.Println("2.修改书籍信息")
	fmt.Println("3.展示书籍信息")
	fmt.Println("4.退出")
}
func main() {
	for{
		showMenu()
		//获取用户输入的值
		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			addBook()
		case 2:
			updateBook()
		case 3:
			showBooks()
		case 4:
			os.Exit(0)

		}
	}
}

func showBooks() {
	if len(allBooks)==0{
		fmt.Println("您的系统中还没有书籍哦Q！")
		return
	}
	for _,b := range allBooks {
		fmt.Printf("<<%s>> 作者:%s 价格:%.2f 是否上架:%t\n",b.title,b.author,b.price,b.publish)
	}

}

func updateBook() {
	book := userInput()
	for i,b := range allBooks {
		if b.title==book.title{
			allBooks[i]=book
			fmt.Printf("<<%s>> 更新成功！\n",book.title)
			return
		}
	}
	fmt.Printf("<<%s>> 不存在！\n",book.title)
}

func addBook() {

	book := userInput()
	//3.将书籍信息添加到books切片中
	//在添加书籍之前可以首先书籍是否已经存在
	for _,b := range allBooks {
		if b.title==book.title{
			fmt.Printf("<<%s>>这本书已经存在！",book.title)
			return
		}
	}
	allBooks = append(allBooks, book)
	fmt.Printf("添加书籍成功！")
}
