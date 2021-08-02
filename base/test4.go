package main

import "fmt"

//iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，
//const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
const (
	a = iota
	b = 8
	c
	d = iota
)

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

func main() {
	println(a, b, c, d)
	println("i=", i)
	println("j=", j)
	println("k=", k)
	println("l=", l)

	x := 1 >> 3
	println(x)
	var y uint64 = 1<<64 - 1 //不减1或者不用uint就会报错
	println(y)
	var z float64 = 1 << 500 //
	println(z)
	println(z * z)
	println(z * z * z)
	fmt.Println(2 < 1)
	fmt.Println(!(2 < 1))

	x += 100
	println(x)
}
