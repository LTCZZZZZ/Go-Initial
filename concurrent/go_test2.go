// 并发测试
package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		go func() {
			time.Sleep(5 * time.Second)
			fmt.Println("sleep 5s")
		}()
		time.Sleep(2 * time.Second)
		fmt.Println("sleep 2s")
	}()

	//time.Sleep(8 * time.Second)
	select {}
}
