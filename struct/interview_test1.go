// 容易犯错的面试题
package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	for _, stu := range stus {
		fmt.Println(stu)
		fmt.Printf("%p\n", &stu) // 每个stu地址相同，大概是因为每次取出值来存入同一个地址
		m[stu.name] = &stu
		fmt.Println(m)
	}
	fmt.Println(m) // 三个地址竟然相同
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
	// 最终输出
	// pprof.cn => 博客
	// 测试 => 博客
	// 博客 => 博客

	// 使用索引，则获取到的是真实值的地址
	for i, stu := range stus {
		//fmt.Println(stus[i])
		fmt.Printf("%p\n", &stus[i]) // 每个stus[i]地址不同，
		m[stu.name] = &stus[i]
		//fmt.Println(m)
	}
	fmt.Println(m) // 三个地址竟然相同
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
	// 最终输出，并注意，map每次输出的顺序不一定相同
	// 测试 => 测试
	// 博客 => 博客
	// pprof.cn => pprof.cn
}
