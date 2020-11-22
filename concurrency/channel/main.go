package channel


import (
"fmt"
"math/rand"
)

//使用goroutine实现一个简单的生产者消费者模型

//生产者：产生随机数 math/rand

//消费者：计算每个随机数的每个数位的数字之和

//1个生产者  20个消费者
var itemChan chan *item
var resultChan chan * result

type item struct{
	id int64
	num int64
}

type result struct {
	item *item
	sum int64
}

//生产者
func producer(ch chan *item)  {
	//生成随机数 并放入通道中
	var id int64
	for {
		id++
		num := rand.Int63()
		tmp := &item{
			id: id,
			num: num,
		}

		//将该结构体放入通道中
		ch<-tmp
	}
	close(ch)
}
//根据传入的数字求各个数字的位数之和
func calc(num int64) int64{
	var sum int64
	for num>0{
		sum=sum+num%10
		num/=10
	}
	return sum
}

//消费者
func consumer(ch chan *item,resultChan chan *result){
	for tmp:= range ch {

		ret := calc(tmp.num)
		retObj:=&result{
			item: tmp,
			sum: ret,
		}

		resultChan <- retObj

	}
}

func startWorker(n int,itemChan chan *item,resultChan chan *result){
	for i:=0;i<n;i++{
		go consumer(itemChan,resultChan)
	}
}
/*
func closeResult(doneChan chan struct{},resultChan chan *result)  {
	for i:=0;i<10;i++{
		<-doneChan
	}
	close(doneChan)
	close(resultChan)
}*/

func printres(resultChan chan *result)  {
	for r := range resultChan {
		fmt.Printf("id:%v num:%v sum:%v\n",r.item.id,r.item.num,r.sum)

	}
}

func main() {
	itemChan =make(chan *item,100)
	resultChan=make(chan *result,100)
	go producer(itemChan)
	startWorker(10,itemChan,resultChan)
	printres(resultChan)
	//给rand加随机数种子实现每一次执行都能产生真正的随机数
	/*rand.Seed(time.Now().UnixNano())
	ret := rand.Int63()
	fmt.Println(ret)
	intn := rand.Intn(101)
	fmt.Println(intn)*/
}