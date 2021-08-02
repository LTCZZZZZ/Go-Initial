package main

import "fmt"

func f0() int {
	i := 5
	defer func() {
		i++
		fmt.Println(i)
	}()
	fmt.Println(i)
	return i - 3
}

// 在 Go 语言的 func 声明中如果返回值变量显示声明，也就是 func foo() (ret int) {} 的时候，rval 就是 ret。
// 否则rval会新建一个变量来储存返回值，就像f0函数一样
func f1() (result int) {
	defer func() {
		result++
	}()
	return 10
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main() {
	println(f0())
	println(f1())
	println(f2())
	println(f3())
}
