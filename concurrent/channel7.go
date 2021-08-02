package main

import (
	"fmt"
	"time"
)

func main()  {
	timerChan := make(chan time.Time)
	fmt.Println(time.Now())
	go func() {
		time.Sleep(2 * time.Second)
		timerChan <- time.Now() // send time on timerChan
	}()
	// Do something else; when ready, receive.
	// Receive will block until timerChan delivers.
	// Value sent is other goroutine's completion time.
	completedAt := <-timerChan
	fmt.Println(completedAt)
}
