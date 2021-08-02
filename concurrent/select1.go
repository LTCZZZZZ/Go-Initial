package main

import (
	"fmt"
	"time"
)

func main()  {
	var ch1, ch2 chan int
	ch1 = make(chan int)
	ch2 = make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 2
	}()

	time.Sleep(1500 * time.Millisecond)

	// 如果有同时多个case去处理,比如同时有多个channel可以接收数据，那么Go会伪随机的选择一个case处理(pseudo-random)。
	// 如果没有case需要处理，则会选择default去处理，如果default case存在的情况下。如果没有default case，则select语句会阻塞，直到某个case需要处理。
	select {
	case v := <-ch1:
		fmt.Println("channel 1 sends", v)
	case v := <-ch2:
		fmt.Println("channel 2 sends", v)
	default: // optional
		fmt.Println("neither channel was ready")
	}
}
