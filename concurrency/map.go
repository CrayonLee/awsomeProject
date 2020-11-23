package main

import (
	"fmt"
	"strconv"
	"sync"
)

//并发安全的map

var m=sync.Map{}

func main() {
	wg:=sync.WaitGroup{}
	wg.Add(20)
	for i:=0;i<20;i++{
		go func(n int) {
			key:=strconv.Itoa(n)
			m.Store(key,n)
			value,_:=m.Load(key)
			fmt.Printf("k:=%v,v:=%v\n",key,value)
			wg.Done()

		}(i)

	}
	wg.Wait()
}