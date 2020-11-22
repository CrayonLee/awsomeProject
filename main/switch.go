package main

func main() {
	a:=[]byte{2,2}
	b:=[]byte{3,2,3}
	i:=Compare(a,b)
	print(i)
}

func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}

	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1

	}
	return 0
}
