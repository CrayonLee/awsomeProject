package main

func main() {
	//var a [3]int
	//a:=[3]int{1,2,3}
	//a:=[...]int{1,2,3}
	//for i, i2 := range a {
	//	print(i,i2)
	//}
	b:=[3][2]int{[2]int{1,2},[2]int{3,4},[2]int{5,6}}
	for _,v := range b {
		for _,v2 := range v {
			print(v2)
		}
	}
}
