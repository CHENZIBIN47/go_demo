package design

import "fmt"

type API interface {
	Say(name string) string
}

func NewApi(i int) API{
	if i == 1{
		return &aAPI{}
	}else if i ==2 {
		return &bAPI{}
	}
	return nil
}

type aAPI struct {

}
type bAPI struct {
}

func (a *aAPI)Say(name string) string  {

	return fmt.Sprintf("aapi,%s",name)
}

func (b *bAPI)Say(name string)string{
	return fmt.Sprintf("bapi,%s",name)
}
