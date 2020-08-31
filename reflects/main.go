package main

import (
	"fmt"
	"myGo/db"
	"reflect"
	"strings"
)

type Student struct {

	Name string `json:"name"`
	Address string `json:"address"`

}
type Customer struct {
	CustId int64 `sql:"cust_id"`
	CommonRegionId  int64 `sql:"common_region_id"`
	PartyId int64 `sql:"party_id"`
	CustNumber string `sql:"cust_number"`
	IsRealName int64 `sql:"is_realname"`
	CustAddr string `sql:"cust_addr"`
	CustType string `sql:"cust_type"`
	GroupCustId int64 `sql:"group_cust_id"`
	StatusCd string `sql:"status_cd"`
	RegionId int64 `sql:"region_id"`
	DistributorId int64 `sql:"distributor_id"`
	CustName string `sql:"cust_name"`

}



func (stu Student)Plain(str string){
	fmt.Println(str)
}

func (stu Student)Read(n int)  {
	fmt.Println(n)
}

func Save(it interface{})bool{
	sql := "insert into "
	typeOf := reflect.TypeOf(it)
	valueOf := reflect.ValueOf(it)
	fields := typeOf.NumField()
	args := make([]interface{},0)
	if typeOf.Kind() != reflect.Struct {
		panic("传进来的类型不是结构体")
	}
	sql += typeOf.Name()+"("
	for i := 0; i < fields; i++ {
		fieldByIndex := typeOf.Field(i)
		value := valueOf.Field(i)
		field := fieldByIndex.Tag.Get("sql")
		sql += field + ","
		//fmt.Println(value.Type())
		//switch value.Interface().(type) {
		//case int64:
		//	args = append(args,value.Int())
		//case string:
		//	args = append(args,value.String())
		//}
		args = append(args,value.Interface())

	}
	sql  = strings.TrimRight(sql,",") + ")"
	sql += "values("
	for j := 0; j < fields; j++ {
		sql += "?,"
	}
	sql  = strings.TrimRight(sql,",") + ")"
	db := db.GetConnect(db.USERNAME, db.PASSWORD, db.NETWORK, db.HOST, db.PORT, db.DATABASE)

	_, err := db.Exec(sql, args...)
	if err != nil {
		panic(err.Error())
	}
	//fmt.Println(result.LastInsertId())
	return true
}
type SqlType string
const (
	SELECT SqlType = "select"
	INSERT SqlType = "insert"
	UPDATE SqlType = "update"
	DELETE SqlType = "delete"

)

func GetSql(i interface{},tableName string, params map[string]interface{},sqlType SqlType){
	typeOf := reflect.TypeOf(i)
	if typeOf.Kind() != reflect.Struct{
		panic("传进来的类型不是结构体")
	}
	sql := ""
	switch sqlType {
	case SELECT:
		sql = "select "
	case INSERT:
		sql = "insert into "


	}
	fmt.Println(sql)
}

func Call(i interface{},methodName string,arg ...interface{})([]reflect.Value,error){
	value := reflect.ValueOf(i)
	//if value.Kind() != reflect.Struct{
	//	return nil,errors.New("类型不对")
	//}
	method := value.MethodByName(methodName) //如果方法是私有的(方法名小写开头的 这边返回的method为nil 将会panic)
	in := make([]reflect.Value,len(arg))
	for k, v := range arg {
		in[k] = reflect.ValueOf(v)
	}
	fmt.Println(in)
	return method.Call(in),nil
}
func main() {
	//customer := Customer{-72,4601,1801402572,"2898000001240067",0,"福建省厦门市","1100",1,
	//	"1000",4601,-1,"litterguo"}
	//
	//Save(customer)
	//student := Student{"chan","123"}
	////Call(student,"Plain","a")
	//Save(student)


	s := "CMDB-process-protocol"
	split := strings.Split(s, "-")
	fmt.Println(len(split))
	for _, v := range split {
		fmt.Println(v)
	}

}
