package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

/*

生产者消费者模型

生产者每秒生产一件商品，并通知物流公司取货
物流公司将商品运输到商店
消费者阻塞等待从商店消费
消费10次主协程结束
*/

type Product1 struct {
	name string
}

func main6() {

	chStorage := make(chan Product1, 5)
	chShop := make(chan Product1, 5)
	wg := &sync.WaitGroup{}
	go Producer(chStorage)

	go Logistics(chStorage, chShop)
	wg.Add(10)
	go Customer(chShop, wg)
	wg.Wait()
}

//消费者
func Customer(chShop <-chan Product1, sywg *sync.WaitGroup) {
	var count int
	for {

		product := <-chShop
		time.Sleep(time.Microsecond)
		fmt.Println("消费了", product.name)
		count++
		fmt.Println("执行了" + strconv.Itoa(count) + "次")
		sywg.Done()

	}
}

//物流
func Logistics(chStorage <-chan Product1, chShop chan<- Product1) {

	for {
		products := <-chStorage
		chShop <- products
	}

}

//生产者
func Producer(chStorage chan<- Product1) {

	for {
		product := Product1{name: "商品" + strconv.Itoa(time.Now().Second())}
		//time.Sleep(time.Second)
		chStorage <- product
	}

}
