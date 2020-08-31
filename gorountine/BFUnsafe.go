package main

import (
	"fmt"
	"sync"
)
//协程同时频繁修改一个数据，会出现并发不安全
func main2() {
	var money = 2000
	var wg sync.WaitGroup
	for i:=0;i<10;i++{
		wg.Add(1)
		go func() {
			for j:=0;j<1000;j++{
				money+=1
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(money)

}



//增加同步锁
func main() {
	//声明同步锁
	mt := sync.Mutex{}
	var money = 2000
	var wg sync.WaitGroup
	for i:=0;i<10;i++{
		wg.Add(1)
		go func() {
			mt.Lock()//抢锁
			for j:=0;j<10000;j++{
				money+=1
			}
			mt.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(money)

}
