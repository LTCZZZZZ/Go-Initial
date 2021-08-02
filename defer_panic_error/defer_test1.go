// defer 是先进后出

package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println(time.Now()) // 首先defer后必须是一个函数，其次此函数的参数（这里是time.Now()）会在这一行立即计算
	time.Sleep(1 * time.Second)
	fmt.Println(time.Now()) // 最后的结果就是后打印的时间比前面还早一秒

	var whatever [5]struct{}
	for i := range whatever {
		defer func() { fmt.Println(i) }() // 输出全为4
	}

}
