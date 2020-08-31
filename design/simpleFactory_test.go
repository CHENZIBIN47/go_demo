package design

import (
	"fmt"
	"testing"
)

func TestSimpleFactory(t *testing.T){

	newApi := NewApi(2)
	say := newApi.Say("hello")
	fmt.Println(say)
}
