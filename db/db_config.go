package db

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type DataBase struct {

	DB *sql.DB
}

const (
	DB_ADDRESS string = "crm_conf/Aa123456@192.168.53.159:28011/CONFDB"
)
func GetConn(dbName string) DataBase{
	var username string
	var password string
	var host string
	var port int
	var database string
	split := strings.Split(dbName, "/")
	username = split[0]
	fmt.Println(split,reflect.TypeOf(split))
	//for _, value := range split {
	//	if strings.Contains(value,"@"){
	//		i := strings.Split(value, "@")
	//		fmt.Println(i)
	//	}
	//}
	str := strings.Split(split[1], "@")
	password = str[0]
	hostPorts := strings.Split(str[1], ":")
	host = hostPorts[0]
	port, _ = strconv.Atoi(hostPorts[1])
	database = split[2]
	fmt.Println(username)
	fmt.Println(password)
	fmt.Println(host)
	fmt.Println(port)
	fmt.Println(database)

	DB,err := sql.Open("mysql",dbName)
	base := DataBase{DB: DB}
	if err != nil {
		fmt.Println("connect failed....")
		return base
	}
	return base
}

func (db *DataBase)Execute(sql *string)([]map[string]interface{}){

	return nil
}

type  CustomerOrder struct {
	CustOrderId int `sql:"cust_order_id"`
	CustOrderNbr string `sql:"cust_number"`
	CustOrderType string `sql:"cust_order_type"`
	CreateDate time.Time `sql:"create_date"`
}

func Save(stu interface{},db *sql.DB) {
	//bufferString := bytes.NewBufferString("insert into ")
	sqlString := "insert into "
	types := reflect.TypeOf(stu)
	values := reflect.ValueOf(stu)
	fields := make([]string, 0)
	fieldValues := make([]interface{}, 0)
	if types.Kind() != reflect.Struct {

	}
	//bufferString.WriteString(types.Name()+"(")
	sqlString = sqlString + types.Name() + "("
	for i := 0; i < types.NumField(); i++ {
		field := types.Field(i)
		fields = append(fields, field.Tag.Get("sql"))
		sqlString += field.Tag.Get("sql") + ","

	}
	sqlString = strings.TrimRight(sqlString, ",") + ")"
	sqlString += "values("
	for i := 0; i < values.NumField(); i++ {
		value := values.Field(i)

		switch value.Interface().(type) {
		case int:
			fmt.Println("is -----int-----")
			fmt.Println(value.Interface().(int))
		case string:
			fmt.Println("is -----string-----")
			fmt.Println(value.Interface().(string))
		case time.Time:
			fmt.Println("is -----time------")
			fmt.Println(value.Interface().(time.Time))
		}
		fieldValues = append(fieldValues, value.Interface())
	}
	fmt.Println(fields)
	fmt.Println(fieldValues)
	fmt.Println(sqlString)

	db.Prepare(sqlString)



}




//var commonInitialisms = []string{"ACL", "API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SQL", "SSH", "TCP", "TLS", "TTL", "UDP", "UI", "UID", "UUID", "URI", "URL", "UTF8", "VM", "XML", "XMPP", "XSRF", "XSS"}
//var commonInitialismsReplacer *strings.Replacer
//var uncommonInitialismsReplacer *strings.Replacer
//
//
//func UnMarshal(name string) string {
//	const (
//		lower = false
//		upper = true
//	)
//
//	if name == "" {
//		return ""
//	}
//
//	var (
//		value                                    = commonInitialismsReplacer.Replace(name)
//		buf                                      = bytes.NewBufferString("")
//		lastCase, currCase, nextCase, nextNumber bool
//	)
//
//	for i, v := range value[:len(value)-1] {
//		nextCase = bool(value[i+1] >= 'A' && value[i+1] <= 'Z')
//		nextNumber = bool(value[i+1] >= '0' && value[i+1] <= '9')
//
//		if i > 0 {
//			if currCase == upper {
//				if lastCase == upper && (nextCase == upper || nextNumber == upper) {
//					buf.WriteRune(v)
//				} else {
//					if value[i-1] != '_' && value[i+1] != '_' {
//						buf.WriteRune('_')
//					}
//					buf.WriteRune(v)
//				}
//			} else {
//				buf.WriteRune(v)
//				if i == len(value)-2 && (nextCase == upper && nextNumber == lower) {
//					buf.WriteRune('_')
//				}
//			}
//		} else {
//			currCase = upper
//			buf.WriteRune(v)
//		}
//		lastCase = currCase
//		currCase = nextCase
//	}
//
//	buf.WriteByte(value[len(value)-1])
//
//	s := strings.ToLower(buf.String())
//	return s
//}

