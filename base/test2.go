package main

import (
	"fmt"
	"io"
	"os"
)

//全局变量是允许声明但不使用，但函数体内的变量不行
var (
	pi = 3.14
	s1 = "China"
)

func main() {
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(url)
	fmt.Println(target_url)

	// 另外一个实例，%d 表示整型
	const name, age = "Kim", 22
	// := 声明并赋值，系统自动推断类型
	s := fmt.Sprintf("%s is %d years old.\n", name, age)
	io.WriteString(os.Stdout, s) // 简单起见，忽略一些错误

	a := 0 //单纯地给 a 赋值也是不够的，这个值必须被使用
	//这种不带声明格式的只能在函数体中出现, 且等号左边必须带有至少一个未初始化的变量
	b, c, a := 1, 2, 3
	fmt.Println(b, c, a)
	var e, f int
	e = b
	fmt.Println(e, f, pi, s1)

}
