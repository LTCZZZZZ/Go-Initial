package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Printf("sum:")
	fmt.Printf("%#v\n", sum)
	fmt.Println("behind channel pro")
	c <- sum // 把 sum 发送到通道 c
	fmt.Println("after channel pro")
}

// 通道不带缓冲，表示是同步的，只能向通道 c 发送一个数据，只要这个数据没被接收则此goroutine阻塞，其他往此通道内发数据的操作同样阻塞
func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	fmt.Println("go [0,3]")
	go sum(s[:len(s)/2], c) //a

	time.Sleep(1000 * time.Millisecond)

	//c <- 1000  // 加上这一行就会deadlock，因为这里阻塞了运行不到下面

	fmt.Println("go [3,6]")
	go sum(s[len(s)/2:], c) //b
	fmt.Println("go2 [0,3]")
	go sum(s[:len(s)/2], c) //c
	fmt.Println("go2 [3,6]")
	go sum(s[len(s)/2:], c) //d

	/*
	   a b c d和main一起争夺cpu的，他们的执行顺序完全无序，甚至里面不同的语句都相互穿插
	   但无缓冲的等待是同步的，所以接下来a b c d都会执行，一直执行到c <- sum后，开始同步阻塞
	   因此after channel pro是打印不出来的, 等要打印after channel pro的时候，main就结束了
	*/

	fmt.Println("go3 start waiting...")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("go3 waited 1000 ms")

	//因为a b c d都在管道门口等着，这里度一个，a b c d就继续一个，这个结果的顺序可能是acbd
	aa := <-c
	bb := <-c
	fmt.Println(aa)
	fmt.Println(bb)
	time.Sleep(1000 * time.Millisecond)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

}
