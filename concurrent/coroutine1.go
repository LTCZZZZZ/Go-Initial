package main

import (
	"fmt"
	"time"
)

func printNumber(from, to int, c chan int) {
	for x := from; x <= to; x++ {
		fmt.Printf("%d\n", x)
		time.Sleep(1 * time.Second)
	}
	c <- 0
}
func main() {
	c := make(chan int, 3)
	go printNumber(1, 5, c)
	go printNumber(4, 6, c)
	_ = <-c
	_ = <-c
	_ = <-c // fatal error: all goroutines are asleep - deadlock!
}
