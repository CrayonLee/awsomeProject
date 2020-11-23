package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
	chan<-int 是一个只写单向通道
	<-chan int是一个只读单向通道
 */

//生成一个int64的随机整数，并发送至jobChan中
func generate(jobChan chan<-int64)  {
	r:=rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<24;i++{
		int63 := r.Int63()
		fmt.Println(int63)
		jobChan<-int63
	}
	close(jobChan)
}

//从jobChan中读出int64型的整数值并发送值resultChan中
func cal(jobChan <-chan int64,resultChan chan<-int64)  {

	var sum int64 =0
	num:=<-jobChan
	for num>0{
		last:=num%10
		sum+=last
		num/=10
	}
	resultChan<-sum
}

func main()  {


	jobChan,resultChan:=make(chan int64,100),make(chan int64,100)
	for i:=0;i<24;i++{
		go cal(jobChan,resultChan)
	}
	go generate(jobChan)
	for i:=0;i<24;i++{
		result:=<-resultChan
		fmt.Println(result)
	}
	close(resultChan)


}