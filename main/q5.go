package main

func average(num [2]float64) (avg float64) {
	sum:=0.0
	switch len(num) {
	case 0:
		avg=0
	default:
		for _,v := range num {
			sum+=v
		}
		avg=sum/float64(len(num))
	}
	return
}

func main() {
	n:=[...]float64{0.0,1.0}
	var avg = average(n)
	print(avg)
}
