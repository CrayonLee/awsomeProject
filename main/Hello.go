package main

func main() {
	/*a1:= "Hello World"
	a := []byte(a1)
	for i,j :=0,len(a)-1;i<j;i,j=i+1,j-1{
		a[i],a[j]=a[j],a[i]
	}

	for i:=0;i<len(a);i++{
		print(a[i])
	}*/
	J:	for j:=0;j<5;j++  {
			for i:=0;i<10;i++{
				if i>5 {
					break J
				}
				println(i)
			}
		}
}
