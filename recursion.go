// 函数内递归，重点参见下面注释
package main

import (
	"fmt"
	"time"
)

func main() {
	Function1(15)
	fmt.Println(Add(10))
	fmt.Println(Factorial(5))

	start1 := time.Now()
	fmt.Println(Fibonacci(40))
	fmt.Println(time.Since(start1))
	start2 := time.Now()
	fmt.Println(Fibonacci2(40))
	fmt.Println(time.Since(start2))
}

func Function1(n int) int {
	a := 10
	var Function2 func(m int) int // Function2必须先声明，然后才能在其内部递归调用

	Function2 = func(m int) int {
		println(m)
		if m <= a {
			return a
		}
		return Function2(m - 1)
	}

	return Function2(n)
}

func Add(n int) int {
	var Function2 func(m int) int

	Function2 = func(m int) int {
		//println(m)
		if m == 1 {
			return 1
		}
		return m + Function2(m-1)
	}

	return Function2(n)

}

func Factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func Fibonacci(n int) int {
	// 此算法当n >= 50时就很慢很慢了，n=50时用时约38s
	if n == 1 || n == 2 {
		return 1
	}
	return Fibonacci(n-2) + Fibonacci(n-1)
}

func Fibonacci2(n int) int {
	// 相比之下，n=50时此算法仅用时18us(微秒)
	l := []int{1, 1}
	for j := 0; j <= n-3; j++ {
		l = append(l, l[j]+l[j+1])
	}
	//fmt.Println(l)
	return l[n-1]
}
