package service

import "fmt"

type impl struct {

}

func Start(i impl) {
	fmt.Println("start")
}

func Stop(i impl){
	fmt.Println("stop")
}
