package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "10.41.12.37:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	for {

		reply, err := redis.String(c.Do("SET", "username", "nick", "EX", 1231231, "NX"))
		fmt.Println(reply)
		if len(reply) == 0 {
			fmt.Println("Exists")
		} else {
			fmt.Println("First Set")
			fmt.Printf("err %v\n", err)
		}

		if err != nil {
			fmt.Println("redis set failed:", err)
		}
		username, err := redis.String(c.Do("GET", "username"))
		if err != nil {
			fmt.Println("redis get failed:", err)
		} else {
			fmt.Printf("Got username %v \n", username)
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
}
