package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	time.Sleep(-1 * time.Second)
	fmt.Println("Sleep -1 second")

	var mutex sync.Mutex
	fmt.Println("start lock main")
	mutex.Lock()
	fmt.Println("get locked main")
	for i := 1; i <= 3; i++ {
		go func(i int) {
			time.Sleep(time.Duration(i) * time.Millisecond)
			fmt.Println("start lock ", i)
			mutex.Lock()
			fmt.Println("get locked ", i)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("Unlock the lock main")
	mutex.Unlock()
	fmt.Println("get unlocked main")
	time.Sleep(time.Second)
}
