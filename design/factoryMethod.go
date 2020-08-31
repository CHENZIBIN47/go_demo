package design


type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}


type OperatorFactory interface {
	Create() Operator
}

type OperatorBase struct {
	a,b int
}


func (o *OperatorBase)SetA(a int)  {
	o.a = a
}
func (o *OperatorBase)SetB(b int)  {
	o.b = b
}

//PlusOperatorFactory 是 PlusOperator 的工厂类
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

//PlusOperator Operator 的实际加法实现
type PlusOperator struct {
	*OperatorBase
}


//Result 获取结果
func (o PlusOperator) Result() int {
	return o.a + o.b
}
