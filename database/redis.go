package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "r-2ze53f1aebd38384pd.redis.rds.aliyuncs.com:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	if res, err := c.Do("AUTH", "Keleshuxue2017"); err != nil {
		fmt.Println(res, err)
		c.Close()
		return
	}

	if _, err := c.Do("SELECT", 13); err != nil {
		fmt.Println(err)
		c.Close()
	}

	fmt.Println("redis conn success")

	rec1, err := redis.String(c.Do("Get","test:hierarchy1:1000"))
	fmt.Println(rec1, err)

	defer c.Close()
}
