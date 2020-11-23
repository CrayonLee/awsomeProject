package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
)

func main()  {
	http.HandleFunc("/web",search)
	http.HandleFunc("/submit",sub)
	http.HandleFunc("/info",info)

	http.ListenAndServe("127.0.0.1:8080",nil)
}

func info(writer http.ResponseWriter, request *http.Request) {
	data, err := ioutil.ReadFile("web/info.html")
	if err!=nil{
		fmt.Println("read html file failed,err:",err)
	}
	intn := rand.Intn(10)
	dataStr := string(data)
	if intn>5{
		dataStr=strings.Replace(dataStr,"{ooxx}","<li>平凡的世界</li>",1)
	}else{
		dataStr=strings.Replace(dataStr,"{ooxx}","<li>傲慢与偏见</li>",1)
	}
	writer.Write([]byte(dataStr))
}

func search(writer http.ResponseWriter, request *http.Request) {
	data,err:=ioutil.ReadFile("web/form.html")
	if err!=nil{
		fmt.Println("read html file failed,err:",err)
	}

	writer.Write(data)
}

func sub(writer http.ResponseWriter, request *http.Request) {
	print(request.Method)
	request.ParseForm()
	form := request.Form
	username := form.Get("username")
	password := form.Get("password")
	fmt.Printf("%#v",form)
	fmt.Println(username,password)

}


