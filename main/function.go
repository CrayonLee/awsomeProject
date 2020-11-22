package main

import "fmt"

func main() {
	a:= func() {
		println("Hello")
	}
	a()
	fmt.Printf("%T\n", a)
}
