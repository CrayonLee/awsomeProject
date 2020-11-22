package main

import (
	"fmt"
	"os"
)

func showMenu()  {
	fmt.Println("欢迎进入Student Management System！")
	fmt.Println("请输入您的的选择")
	fmt.Println("1.添加学生信息")
	fmt.Println("2.修改学生信息")
	fmt.Println("3.删除学生信息")
	fmt.Println("4.查看所有学生信息")
	fmt.Println("5.退出")
}

func userInput() *Student{
	var (
		name string
		age int
		id int64
		class string
	)
	//1.获取用户输入
	fmt.Println("请根据提示输入相关内容")
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Print("请输入学生编号：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生所属班级：")
	fmt.Scanln(&class)

	//2.创建一本一个学生并将其返回
	return NewStudent(name,age,id,class)
	//fmt.Printf("%v\n",book)

}

func getInputId() int64 {
	var id int64
	fmt.Print("请输入需要被删除用户的id：")
	fmt.Scanln(&id)
	return id
}

func main() {
	stuMgr := NewStuMgr()
	for  {
		showMenu()
		fmt.Print("请输入您的的选择:")
		var option int
		fmt.Scanln(&option)
		fmt.Printf("您的选择是%d\n",option)
		switch option {
		case 1:
			input := userInput()
			stuMgr.AddStudnet(input)
		case 2:
			input := userInput()
			stuMgr.EditStudent(input)
		case 3:
			/*input := userInput()
			stuMgr.deleteStudent(input)*/
			id := getInputId()
			stuMgr.deleteStudent(id)
		case 4:
			stuMgr.showAllStudnets()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("您输入的选项无效")
		}
	}
}
