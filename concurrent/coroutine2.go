package main

import (
	"fmt"
	"runtime"
	"time"
)

func TestFunc(n int, c chan time.Duration) time.Duration {
	time0 := time.Now()

	for i := 0; i < n; i++ {
		a := i
		a = a + 1
	}

	time1 := time.Now()
	c <- time1.Sub(time0)
	return time1.Sub(time0)
}

func main() {
	// 为什么单独运行一个 TestFunc(20000000, c)的时间大约是协程运行2个TestFunc(10000000, c)的两倍？
	// 虽然是协程，但应该也是同一个线程运行(且这里也没有发生系统调用，基本没有阻塞的地方)，
	// 按道理来说时间应该是差不多的，但实际测试并非如此，原因如下：
	// 假如设置runtime.GOMAXPROCS(1)，则协程运行2个TestFunc(10000000, c)的时间也为5ms左右
	time0 := time.Now()
	c := make(chan time.Duration, 3)

	// n=2大约2.5ms
	// 经测试猜测和我电脑的核数或go的设置有关，n<=8时运行时间都差不多，n>8时时间明显增加
	// n=100时可以看到，没有任何一个协程运行时间超过10ms，但总时长超过了35ms
	fmt.Println(runtime.GOMAXPROCS(16)) // 虽然这个值是16，但经测试感觉临界值是8，参考下面的测试数据

	//fmt.Println(runtime.NumCPU())
	// 重要提示：虽然我电脑的runtime.NumCPU()=16，但runtime.GOMAXPROCS()设置更大的值，比如32，
	// 基本就无法提速了，数据如下：
	// 设置n=32，GOMAXPROCS(16)和GOMAXPROCS(32)时，运行时间均为12ms左右。GOMAXPROCS(8)平均约为13ms，不稳定
	// 设置n=100，GOMAXPROCS(16)和GOMAXPROCS(32)时，运行时间均为35ms左右。GOMAXPROCS(8)平均约为36ms，不稳定
	// 猜测分析：和CPU的超线程技术有关

	n := 100
	for i := 0; i < n; i++ {
		go TestFunc(10000000, c)
	}
	for i := 0; i < n; i++ {
		fmt.Println(<-c)
	}

	// 大约5ms
	//go TestFunc(20000000, c)
	//fmt.Println(<-c)

	fmt.Println("总运行时间：", time.Now().Sub(time0))
}
