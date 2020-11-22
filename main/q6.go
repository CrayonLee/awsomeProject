package main

func order(a,b int)(int,int)  {
	if(b>a){
		return b,a
	}
	return a,b
}
