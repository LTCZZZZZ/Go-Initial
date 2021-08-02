// 我们希望能在time-consuming operation 之前就释放锁，而不是等到整个foo返回。
//这有两个办法，一个是根据逻辑，把foo拆分两部分，前半部分需要锁，后半部分不需要锁；

//另一个办法是使用匿名函数：  （暂时没看出来有什么特殊作用，先留存待观）

package main

import "log"
import "time"
import "sync"

var mu sync.Mutex

func lock() {
	mu.Lock()
	log.Printf("lock")
}

func unlock() {
	mu.Unlock()
	log.Printf("unlock")
}

func foo() int {
	lock()

	func() {
		log.Printf("entry inner")
		defer unlock()
		log.Printf("exit inner")
	}()

	time.Sleep(1 * time.Second)
	log.Printf("return")
	return 0
}

func main() {
	r := foo()
	log.Println("r=", r)
}
