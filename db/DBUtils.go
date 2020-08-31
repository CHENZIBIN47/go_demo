package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"

)

const (
	USERNAME = "crm_cust"
	PASSWORD = "Aa123456"
	NETWORK = "tcp"
	HOST = "192.168.53.159"
	PORT = 28015
	DATABASE = "cusdb"
)

var DB *sql.DB

func GetConnect(userName string, password string, netWork string, host string, port int, dataBase string) *sql.DB {

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",userName,password,netWork,host,port,dataBase)
	DB,err := sql.Open("mysql",dsn)
	if err != nil {
		fmt.Println("connect failed....")
		return nil
	}
	DB.SetConnMaxLifetime(100*time.Second)  //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)//设置最大连接数
	DB.SetMaxIdleConns(16) //设置闲置连接数
	return  DB
}
