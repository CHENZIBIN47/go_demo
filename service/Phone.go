package service

import "fmt"

type Phone struct {

}

func (p Phone) Start()  {
	fmt.Println("phone start")
}

func (p Phone) Stop(){
	fmt.Println("phone stop")
}
