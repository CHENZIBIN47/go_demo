//package main
//
//import (
//	"context"
//	"fmt"
//	"myGo/db"
//	"myGo/grpc/proto"
//	"runtime"
//	"sort"
//	"strconv"
//	"sync"
//
//	//"github.com/garyburd/redigo/redis"
//
//)
//
//func main2() {
//
//	//redis.Dial("","")
//	connect := db.GetConnect(db.USERNAME, db.PASSWORD, db.NETWORK, db.HOST, db.PORT, db.DATABASE)
//	row := connect.QueryRow("select * from customer where cust_id=?", 721)
//	customer := new(Cusomer)//这个new出来的是一个指针
//	row.Scan(&customer.CustName)
//	fmt.Println(row)
//	//r := gin.Default()
//	//r.GET("/ping", func(c *gin.Context) {
//	//	c.JSON(200, gin.H{
//	//		"message": "pong",
//	//	})
//	//})
//	//r.Run() // listen and serve on 0.0.0.0:8080
//	//var name string = "chan"
//	//inst := service.AddProdInst(name)
//	//fmt.Println(inst)
//	//service.Break()
//	// var usb service.Usb
//	// //golang 接口
//	// //第一种
//	// //usb = service.Phone{}
//	// //第二种
//	// usb = new (service.Phone)
//	//
//	// //p := service.Phone{}
//	// //c := service.Camera{}
//	// //computer := service.Computer{}
//	//
//	//usb.Start()
//	// //computer.Working(p)
//	// //computer.Working(c)
//	// str :=[]string{"a","b","c"}
//	//for _, value := range str {
//	//
//	//	fmt.Println(value)
//	//
//	//}
//
//
//	//golang 切片（集合）
//	//slice()
//
//
//	//golang map测试
//	mapTest()
//
//	//结构体
//	var p1 Person //定义结构体
//	p1.age  = 11
//	p1.Name = "d"
//	testPerson(&p1)
//	fmt.Println(p1)
//
//	var person1 *Person = new (Person) //定义结构体指针
//	person1.Name = "a"
//	person1.age = 18
//	testPerson(person1)
//	fmt.Println(*person1)
//
//
//	//结构体方法
//	fmt.Println("------------------结构体方法start-------------------")
//	var cust Cusomer
//	cust.CustName = "chena"
//	cust.party(&cust)
//	fmt.Println(cust.CustName)
//	fmt.Println("------------------结构体方法end-------------------")
//
//
//}
//
//
//
//func slice()  {
//	str := make([]string,0)
//	str = append(str, "11","22","33")//向集合中添加元素
//	for key, value := range str {
//		fmt.Println(strconv.Itoa(key)+"..."+value)//key是int类型需要强转为string类型才能进行拼接
//	}
//	fmt.Println(len(str)) //获取切片长度
//
//	str2 := make([]string,10)
//	copy(str2,str)
//	for _, value2 := range str2 {
//		fmt.Println(value2)
//	}
//}
//
//func mapTest() {
//	map1 := make(map[string]string)
//
//	map1["name"] = "chen"
//	map1["age"] = "18"
//	map1["password"] = "123"
//	map1["name"] = "lin"  //key 出现重复会替换原先的
//	for key, value := range map1 {
//		fmt.Println("key:"+key+" value:"+value)
//	}
//	//fmt.Println(map1)
//
//	//对map进行有序排序
//
//	var keys []string
//	for key, _ := range map1 {
//		keys = append(keys,key)
//	}
//	sort.Strings(keys)
//
//	for _, value := range keys {
//
//		fmt.Println(value + ":::v:::"+map1[value])
//		//fmt.Println(map1[value])
//
//	}
//
//	//map切片
//	maps := make([]map[string]string,0) //切片想从第一个元素增加数据 数据写0
//
//	map2 := make(map[string]string)
//
//	map2["name"] = "aa"
//	map2["age"] = "18"
//	map2["password"] = "111"
//	maps = append(maps,map1,map2)
//	fmt.Println("----------------------map切片开始---------------------")
//	for key, value := range maps {
//		fmt.Println(key)
//		fmt.Println(value)
//	}
//	fmt.Println("----------------------map切片结束---------------------")
//}
//
// type Person struct {
// 	Name string
// 	age int
//
// }
////结构体是值拷贝 只能使用指针才能修改原先的值
//func testPerson(person *Person)  {
//	person.Name = "bb"
//	person.age = 28
//	fmt.Println(*person)
//}
//
//type Cusomer struct {
//	CustName string
//}
//
////一般这样写
//func (c *Cusomer) party2()  {
//
//}
//func (cust Cusomer) party(c *Cusomer) {
//	c.CustName = "qaz"
//	cust.CustName = "qwer"
//	fmt.Println(cust.CustName)
//	fmt.Println(c.CustName)
//}
//
//const (
//
//	port = ":8083"
//)
//type server struct{}
//
//func (s *server) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloResponse, error) {
//	return &hello.HelloResponse{Message:in.Value}, nil
//}
//
////func main() {
////	//db.GetConn(db.DB_ADDRESS)
////
////	//customerOrder := db.CustomerOrder{1, "HN2013061600933501", "1000",time.Now()}
////	//db.Save(customerOrder)
////	go func() {
////		conn, err :=  grpc.Dial("localhost:8083",grpc.WithInsecure())
////		if err != nil {
////			log.Fatalf("did not connect: %v", err)
////		}
////		defer conn.Close()
////		c := hello.NewHelloServiceClient(conn)
////		name := "world"
////		response, e := c.SayHello(context.Background(), &hello.HelloRequest{Value: name})
////		if e != nil{
////			log.Fatalf("could not greet: %v", err)
////		}
////
////		log.Printf("hello: %s", response.Message)
////
////	}()
////
////
////	lis, err := net.Listen("tcp", port)
////	if err != nil {
////		log.Fatalf("failed to listen: %v", err)
////	}
////	s := grpc.NewServer()
////	hello.RegisterHelloServiceServer(s,&server{})
////	reflection.Register(s)
////	if err := s.Serve(lis); err != nil {
////		fmt.Println("server success")
////		log.Fatalf("failed to serve: %v", err)
////	}
////
////
////}
//
////func main(){
////
////
////
////
////
////	msg := fmt.Sprintf("这是测试任务")
////	t := &TestPro{
////		msg,
////	}
////	queueExchange := &rabbitmq.QueueExchange{
////		QuName: "test.rabbit",
////		RtKey:  "rabbit.key",
////		ExName: "test.rabbit.mq",
////		ExType: "direct",
////	}
////	mq := rabbitmq.New(queueExchange)
////	mq.RegisterProducer(t)
////	mq.RegisterReceiver(t)
////	mq.Start()
////}
//
//func main() {
//	type MyInt1 int
//	type MyInt2 = int
//	var i int =9
//	var i1 MyInt1 = i
//	var i2 MyInt2 = i
//	fmt.Println(i1,i2)
//}

package main

import (
	"fmt"
	"log"
	"myGo/crontab"
	"time"
)

var data = `
blog: xiaorui.cc
best_authors: ["fengyun","lee","park"]
desc:
  counter: 521
  plist: [3, 4]
`

type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func main() {
	//	t := T{}
	//
	//	err := yaml.Unmarshal([]byte(data), &t)
	//	if err != nil {
	//		log.Fatalf("error: %v", err)
	//	}
	//	fmt.Printf("--- t:\n%v\n\n", t)
	//
	//	d, err := yaml.Marshal(&t)
	//	if err != nil {
	//		log.Fatalf("error: %v", err)
	//	}
	//	fmt.Printf("--- t dump:\n%s\n\n", string(d))
	//
	//	m := make(map[interface{}]interface{})
	//
	//	err = yaml.Unmarshal([]byte(data), &m)
	//	if err != nil {
	//		log.Fatalf("error: %v", err)
	//	}
	//	fmt.Printf("--- m:\n%v\n\n", m)
	//
	//	d, err = yaml.Marshal(&m)
	//	if err != nil {
	//		log.Fatalf("error: %v", err)
	//	}
	//	fmt.Printf("--- m dump:\n%s\n\n", string(d))
	//}

	ctab := crontab.New() // create cron table

	// AddJob and test the errors
	err := ctab.AddJob("* * * * *", myFunc) // on 1st day of month
	if err != nil {
		log.Println(err)
		return
	}

	// MustAddJob is like AddJob but panics on wrong syntax or problems with func/args
	// This aproach is similar to regexp.Compile and regexp.MustCompile from go's standard library,  used for easier initialization on startup
	ctab.MustAddJob("* * * * *", myFunc)   // every minute
	ctab.MustAddJob("0 12 * * *", myFunc3) // noon lauch

	// fn with args
	ctab.MustAddJob("0 0 * * 1,2", myFunc2, "Monday and Tuesday midnight", 123)
	ctab.MustAddJob("*/5 * * * *", myFunc2, "every five min", 0)

	// all your other app code as usual, or put sleep timer for demo
	time.Sleep(10 * time.Minute)
}

func myFunc() {
	fmt.Println("Helo, world")
}

func myFunc3() {
	fmt.Println("Noon!")
}

func myFunc2(s string, n int) {
	fmt.Println("We have params here, string", s, "and number", n)
}
