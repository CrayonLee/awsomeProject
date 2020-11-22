package concurrency

import (
	"fmt"
	"sync"
)

//启动多个goroutine 并实现多个goroutine间的同步
var wg sync.WaitGroup

func Hello(i int)  {
	// goroutine结束就登记-1
	defer wg.Done()
	fmt.Println("Hello Goroutine ",i)

}
func main() {
	for i:=0;i<10;i++ {
		//启动一个goroutine就登记+1
		wg.Add(1)
		go Hello(i)
	}
	//等待所有等级的goroutine都结束
	wg.Wait()
}
