package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	/* 这是我的第一个简单的程序 */
	//time.Sleep(time.Duration(2)*time.Second)
	fmt.Println("Hello, World!" + "123")
	//i := 0
	//for {
	//	println(i)
	//	i ++
	//	time.Sleep(1 * time.Second)
	//	if i >= 3 {
	//		break
	//	}
	//}
	balance := make([]float32, 5)
	fmt.Println(balance)
	balance[3] = 3
	balance[4] = 4.5
	balance = append(balance, 20)
	fmt.Println(balance)
	fmt.Println(3 ^ 2)
	fmt.Println(math.Pow(3, 2))
	fmt.Println(math.Pow(2, 3))

	start := time.Now()
	//time.Sleep(1*time.Second)
	during := time.Since(start)
	fmt.Println(start)
	fmt.Println(during)

	a, b := 1, 2
	println(a, b)
	a, b = b, a
	println(a, b)

	fmt.Println(3 / 2)
	var c float32 = 3 / 2
	fmt.Println(c)
	fmt.Println(float32(3) / 2)
	fmt.Println(3 / float32(2))

	// 下面这两个输出竟然不一样，醉了，之后有空再研究其中的细节吧，python中如果不用demical模块的话1.69 * 1.7输出的也是2.8729999999999998
	// python中要用Decimal('1.7') * Decimal('1.69')，注意传入数据时要用字符串而不能用浮点数，
	//fmt.Println(1.69 * 1.7)
	// 另外，我个人猜想，因为最后一位是8所以才有此问题，如果小数最后很多位都是9，则不会有此问题，否则1.69打印时不可能显示1.69
	m := 1.69
	n := 1.7
	o := 1.69 * 1.7
	u := 1.5
	fmt.Println(m * u)
	fmt.Println(m, n, o, m*n)
}
