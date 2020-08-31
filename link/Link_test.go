package link

import (
	"container/list"
	"fmt"
	"testing"
)

func Test_Link(t *testing.T){
	link := list.New()

	for i :=0; i<=10; i++ {
		link.PushBack(i)
	}
	for p := link.Front(); p != link.Back(); p = p.Next() {
		fmt.Println("Number", p.Value)
	}
}
