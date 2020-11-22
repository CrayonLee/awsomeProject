package concurrency

import "fmt"
func rev(c chan int){
	ret:=<-c
	fmt.Println("接收成功",ret)
}
func main() {
	//这段代码会造成死锁  因为我们在这里创建的是一个无通道缓存 必须等到有人去接收时值才会被发送
	ch:=make(chan int)
	//启动一个新的goroutine用于接收值
	go rev(ch)
	ch <- 10
	fmt.Println("发送成功")

}
