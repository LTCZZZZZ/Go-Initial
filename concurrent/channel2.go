package main

import "fmt"

func main() {
	// 这里我们定义了一个可以存储整数类型的带缓冲通道
	// 缓冲区大小为2
	ch := make(chan int, 5) // 如果是下面的写法的话，这里必须加缓冲区，否则会报错，详见channel4注释

	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 2
	close(ch)
	v, ok := <-ch
	fmt.Println(v, ok)
	v, ok = <-ch
	fmt.Println(v, ok)
	v, ok = <-ch
	fmt.Println(v, ok)
	//ch <- 3
	//fmt.Println(temp)

	// 获取这两个数据
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
}
