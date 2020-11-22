package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID int
	Gender string
	Name string
}

func main() {
	var stu = Student{
		ID: 1,
		Gender: "男",
		Name:"zs",
	}
	//将对象序列化成Json
	v,err:=json.Marshal(stu)
	if err!=nil{
		fmt.Println("Json序列化出错")
		fmt.Println(err)
	}
	fmt.Printf("%v\n",string(v))

	//反序列化
	str:="{\"ID\":1,\"Gender\":\"男\",\"Name\":\"zs\"}"
	var stu2= &Student{}
	json.Unmarshal([]byte(str),stu2)
	fmt.Println(stu2)
	fmt.Printf("%T\n",stu2)
}
