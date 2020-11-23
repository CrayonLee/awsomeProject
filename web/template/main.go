package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Username string
	Password string
	Age int
}
func main() {
	http.HandleFunc("/info",info)
	http.HandleFunc("/info1",info1)
	http.ListenAndServe("127.0.0.1:8080",nil)
}

func info1(writer http.ResponseWriter, request *http.Request) {
	files, err := template.ParseFiles("web/template/info1.html")
	if err!=nil{
		fmt.Println("打开文件失败,err:",err)
		return
	}



	userMap:=map[int]User{
		1:{Username: "test1",Password: "1234",Age: 18},
		2:{Username: "test2",Password: "1234",Age: 18},
		3:{Username: "test3",Password: "1234",Age: 18},
	}
	files.Execute(writer,userMap)
}

func info(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	files, err := template.ParseFiles("web/template/info.html")
	if err!=nil{
		fmt.Println("打开文件失败,err:",err)
		return
	}

	user:=&User{Username: "Alex",Password: "Alex123",Age: 18,}
	files.Execute(writer,user)
}
