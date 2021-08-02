package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 10)

	go fibonacciS(cap(c), c)

	for v := range c {
		fmt.Println("out:", time.Now())
		fmt.Println(v)
	}
}

func fibonacciS(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		c <- x
		fmt.Println("in:", time.Now())
		time.Sleep(500 * time.Millisecond)
		x, y = y, x+y
	}

	close(c)
}
