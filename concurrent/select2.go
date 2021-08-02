// 重点：如果c无缓存的话，除非有额外的goroutine有从c中读数据的请求，否则会一直阻塞，即，就写入不进去
// 即若Channel没有缓存，只有sender和receiver都准备好了后它们的通讯(communication)才会发生(否则会Blocking)。
// 如果设置了缓存，就有可能不发生阻塞， 只有buffer满了后 send才会阻塞， 而只有缓存空了后receive才会阻塞。
// 一个nil channel不会通信。
// channel的 receive支持 multi-valued assignment，v, ok := <-ch，它可以用来检查Channel是否已经被关闭了。
// 参考 https://www.runoob.com/w3cnote/go-channel-intro.html
package main

import (
	"fmt"
	"time"
)

// 函数中一般标明chan的方向防止出错或滥用，在此函数中c表示只写channel，quit表示只读channel
func fibonacci(c chan<- int, quit <-chan int) {
	x, y := 0, 1
	for {
		time.Sleep(500 * time.Millisecond)

		//c <- x
		//fmt.Println("通道c写入了数据")  // 在c无缓存且没有接受者的情况下永远执行不到这里
		// 简而言之，如果c无缓存的话，除非有额外的goroutine有从c中读数据的请求，否则会一直阻塞，即，就写入不进去

		// 如果有同时多个case去处理,比如同时有多个channel可以接收数据，那么Go会伪随机的选择一个case处理(pseudo-random)。
		// 如果没有case需要处理，则会选择default去处理，如果default case存在的情况下。如果没有default case，则select语句会阻塞，直到某个case需要处理。
		select {
		case c <- x:
			fmt.Println(x)
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// 这个程序整体很巧妙，fibonacci本是一个无限循环函数，向c中发送数据，却通过select中quit接收数据的方式终结
// 整个模型类似于生产者-消费者循环，主函数中启动消费者的需求，然后wait，然后启动生产者的模型，
// 生产者不停地生产数据传入通道c(此处c没加缓存，一般实际应用时会有缓存，因为消费者处理数据并进行下次接收可能也需要一定的时间)，
// 消费者一直接收并处理数据，不再需要数据后向quit内传入一个信号，此时生产者端c <- x因为c无缓存无法执行，故而一定执行<-quit，
// 于是终结整个生产者程序。当c有缓存时，由于随机执行的情况，可能会在c的缓存填充到一定程度时停止生产。
func main() {
	c := make(chan int)
	//c := make(chan int, 5)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)  // 如果注释掉fibonacci(c, quit)，此处在循环的第一轮就会阻塞
			time.Sleep(500 * time.Millisecond)  // 虽然这里和上面都sleep了500ms，但实际执行间隔还是大约只有500ms，这也就是协程的好处
		}

		// 上面的代码可以改成这样，如果不加退出条件，即成了无限循环
		//for i := range c {
		//	fmt.Println(i)
		//	// 此处可加上退出循环(即不再接收数据)的条件
		//}

		quit <- 0
	}()

	//time.Sleep(1 * time.Second)
	fibonacci(c, quit)
}
