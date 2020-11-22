package main

import "fmt"

func main() {
	//将字符数组从位置4开始的三个字符开始替换为abc
	str:="asSASA ddd dsjkdsjs dk"
	r:=[]rune(str)
	copy(r[4:4+3],[]rune("abc"))
	fmt.Printf("Before: %s\n", str);
	fmt.Printf("After : %s\n",string(r))
}
