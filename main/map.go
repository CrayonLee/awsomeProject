package main

import "fmt"

func main() {
	monthdays := map[string]int{"Jan": 31, "Feb": 28, "Mar": 31, "Apr": 30, "May": 31, "Jun": 30, "Jul": 31, "Aug": 31, "Sep": 30, "Oct": 31, "Nov": 30, "Dec": 31}
	fmt.Printf("%d\n",monthdays["Dec"])

	//当只需声明一个map时 使用make
	//monthdays := make(map[string]int)
	daysOfYear := 0
	for _,day := range monthdays {
		daysOfYear+=day
	}
	fmt.Printf("Numbers of days in a year: %d\n", daysOfYear)

	//var value int
	//var present bool
	//value,present=monthdays["Jan"]
	value,present := monthdays["Jan"]
	print(value,present)
}
