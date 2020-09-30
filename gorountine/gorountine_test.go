package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_limit(t *testing.T) {
	var wg = sync.WaitGroup{}
	number := 8765
	g := New(1000)
	for i := 0; i < number; i++ {
		wg.Add(1)
		value := i
		goFunc := func() {
			// 做一些业务逻辑处理
			fmt.Printf("go func: %d\n", value)
			time.Sleep(time.Second * 2)
			wg.Done()
		}
		g.Run(goFunc)
	}
	wg.Wait()
}

type Glimit struct {
	n int
	c chan struct{}
}

func New(n int) *Glimit {
	return &Glimit{
		n: n,
		c: make(chan struct{}, n),
	}
}

func (g *Glimit) Run(f func()) {
	g.c <- struct{}{}
	go func() {
		f()
		<-g.c
	}()
}
