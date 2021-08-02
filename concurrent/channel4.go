// 终于搞明白了，channel有无缓冲的本质在于是否需要接收方
//（当无缓冲时接收方不能是自己，只能是除自己以外的goroutine）

package main

import (
	"fmt"
	"time"
)

func say1(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, (i+1)*100)
	}
}
func say2(s string, ch chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Println(s, (i+1)*150)
	}
	// 我们引入一个信道，默认的，信道的存消息和取消息都是阻塞的，
	// 在 goroutine 中执行完成后给信道一个值 0，则主函数会一直等待信道中的值，一旦信道有值，主函数才会结束。
	fmt.Println("say2", <-ch)
	ch <- 20 // 令我惊讶的是这一行是可省略的
	ch <- 10
	close(ch) // 经测试发现，close后如果读取通道数据的话，不论多少次，会一直读取到0
	// 此外，关闭通道并不会丢失里面的数据，只是让读取通道数据的时候不会读完之后一直阻塞等待新数据写入————重要！！
}

func main() {
	ch := make(chan int)
	//ch := make(chan int, 10)
	say1("hello")
	go say2("world", ch)
	ch <- 200
	//ch <- 100
	fmt.Println(<-ch) // 这一行，使主函数等待ch中的值
	fmt.Println(<-ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
