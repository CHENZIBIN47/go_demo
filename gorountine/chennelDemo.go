package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GetRandom(start,end int) int{
	rand.New(rand.NewSource(time.Now().UnixNano()))

	return start + rand.Intn(end - start+1)
}

func IsEven(n int) bool{

	return n%2 == 0
}

func IsSx(n int) bool{

	a :=n/100
	b :=n%100/10
	c := n%10
	return (a*a*a + b*b*b + c*c*c) ==n
}

func main() {

	ints := make(chan int, 5)
	even := make(chan bool, 5)
	quit := make(chan string)

	go func() {
		ticker := time.NewTicker(1 * time.Microsecond)
		for{
			random := GetRandom(100, 999)
			fmt.Println(random)
			ints<- random
			<- ticker.C
		}

	}()

	go func() {

		for x := range ints {
			isEven := IsEven(x)

			even<- isEven
			if isEven {
				fmt.Println("o")
			}else {
				fmt.Println("j")
			}
			if IsSx(x) {
				quit <- "quit"
			}
		}
	}()



	go func() {
		var oddCount, evenCount int
		for x := range even {
			if x {
				evenCount++
			}else {
				oddCount++
			}
		}

	}()

	<-quit
}
