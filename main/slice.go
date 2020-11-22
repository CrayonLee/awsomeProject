package main

func main() {
	var a =[...]int{1,2,3,4,5,6,7}
	var s = make([]int,6)
	n1 := copy(s,a[0:])
	n2 := copy(s,s[2:])
	for _,v := range s {
		print(v)
	}
	println()
	print(n1)
	print(n2)
}