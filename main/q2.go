package main

import "fmt"

//FizzBuzz
func main() {
	var arr [100]int
	for i:=0;i<len(arr);i++{
		arr[i]=i+1
	}

	for i := 0; i < len(arr); i++ {
		f:=false
		if(arr[i]%3==0){
			fmt.Printf("%s","Fizz")
			f=true
		}
		if(arr[i]%5==0){
			fmt.Printf("%s","Buzz")
			f=true
		}
		if !f{
			fmt.Printf("%d",arr[i])
		}
		fmt.Println()

	}

/*	const (
		FIZZ=3
		BUZZ=5
	)
	var flag bool
	for i := 1; i <100 ; i++ {
		flag =false
		if(i%FIZZ==0){
			fmt.Printf("Fizz")
			flag=true
		}
		if(i%BUZZ==0){
			fmt.Printf("Buzz")
			flag=true
		}
		if !flag{
			fmt.Printf("%v",i)
		}
		fmt.Println()
	}*/
}
