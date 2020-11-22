package main

import "fmt"

func main() {
	//逆转字符串
	var s string ="Foo Bar"
	a:=[]rune(s)
	for i,j:=0,len(a)-1;i<j;i,j=i+1,j-1{
		a[i],a[j]=a[j],a[i]
	}
	fmt.Printf("%s\n",string(a))
}
