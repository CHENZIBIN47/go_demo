package service

import "fmt"

type Camera struct {

}
func (c Camera) Start()  {
	fmt.Println("camera start")
}
func (c Camera) Stop()  {
	fmt.Println("camera stop")
}
