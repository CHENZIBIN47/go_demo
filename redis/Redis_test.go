package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"testing"
)

func Test_Redis(t *testing.T){

	c, err := redis.Dial("tcp", "192.168.137.129:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

}
