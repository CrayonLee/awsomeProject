package main

import "fmt"

func main() {
	list:= []string{"a","b","c"}
	for k,v:=range list{
		println(k,v)
	}

	for pos,char:=range "abx"{
		fmt.Printf("character %c starts at byte position %d\n",char,pos )
	}
}

