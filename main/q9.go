package main

import "fmt"

func main() {
	prtthem(1,2,3)
}
func prtthem(nums ...int)  {
	for _,v := range nums {
		fmt.Printf("%d ",v)
	}
}
