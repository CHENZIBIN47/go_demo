package main

import (
	"fmt"
)

func main8() {

	//done := make(chan int,1)
	//
	//go func(){
	//	fmt.Println("你好, 世界")
	//	//done <- 1
	//}()
	//<- done

	ch1 := make(chan int, 5)
	ch2 := make(chan int, 3)
	ch3 := make(chan string, 2)

	go AddChan1(ch1)
	go AddChan2(ch2)

	//go ReadChan(ch3)
LABEL:
	for {
		select {
		case x := <-ch1:
			fmt.Println("ch1", x)
		case x := <-ch2:
			fmt.Println("ch2", x)
		case ch3 <- "name":
			fmt.Println("ch3")
		default:
			fmt.Println("break")
			break LABEL
		}
	}

	fmt.Println("main over")

}

func AddChan1(ch chan<- int) {

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5

}

func AddChan2(ch chan<- int) {
	ch <- 7
	ch <- 8
	ch <- 9

}

func ReadChan(ch <-chan string) {
	for v := range ch {
		fmt.Println(v)
	}
}

type Product struct {
	Name string
}
