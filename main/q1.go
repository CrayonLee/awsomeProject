package main

import "fmt"

func main() {
	//1
	/*for i:=0;i<10;i++{
		fmt.Println(i+1)
	}*/

	//使用goto修改循环
	/*i := 0
Loop:
	if i < 10 {
		fmt.Printf("%d\n", i+1)
		i++
		goto Loop
	}*/

	arr:=[...]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Printf("%v",arr)

}
