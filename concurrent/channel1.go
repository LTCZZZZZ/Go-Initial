package main

import (
	"fmt"
	"time"
)

func sum0(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
	time.Sleep(2 * time.Second)
	fmt.Println(sum) // 这一行不会打印，因为主进程已经执行完了，除非加上c <- 100，参见channel4
	time.Sleep(1 * time.Second)
	c <- 100
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum0(s[:len(s)/2], c)
	fmt.Println("1 start")
	go sum0(s[len(s)/2:], c)
	fmt.Println("2 start")
	//x, y := <-c, <-c // 从通道 c 中接收
	x := <-c
	y := <-c
	z := <-c

	fmt.Println(x, y, x+y, z)
}
