package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
var wg sync.WaitGroup
//使用goroutine和channel实现一个int64随机数个位数和的程序

//计算一个64位随机数的个位数之和

func randNumber(x int64) int64{
	var sum int64 =0
	for x>0{
		a:=x%10
		x/=10
		sum+=a
	}
	return sum
}

//生成int64的随机数放入通道jobChan中
func createRand(jobChan chan<- int64){
	for{
		num:=rand.Int63()
		jobChan <- num
		time.Sleep(1)
	}
}

//从jobChan中读取数据，然后计算随机数的各位数之和，将结果发送到resultChan
func readRand(jobChan <-chan int64,resultChan chan<- int64)  {
	for{
		value:=<-jobChan
		num:=randNumber(value)
		resultChan<-num
		fmt.Println(value,num)
	}
}

func main(){
	var jobChan = make(chan int64,100)
	var resultChan = make(chan int64,100)
	wg.Add(25)
	go createRand(jobChan)
	for i:=0;i<24;i++{
		go readRand(jobChan,resultChan)
	}
	for val := range resultChan {
		fmt.Println(val)

	}

}