package main

import (
	"fmt"
	"sync"
)

//等待组：每增加一个子协程，就向等待组中+1,每结束一个协程，就从等待组中-1.主协程会阻塞等待直到组中的协程等于0为止
//这种方式可以令主协程结束在最后一个子协程结束的时间上。
func main(){


	var wg sync.WaitGroup
	fmt.Println(wg)
	//向等待组中添加一个协程
	wg.Add(1)

	//从等待组中减掉一个协程
	wg.Done()
	//阻塞等待至等待组中的协程归零
	wg.Wait()
}
