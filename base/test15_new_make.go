// 在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。
//而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。
//要分配内存，就引出来今天的new和make。 Go语言中new和make是内建的两个函数，主要用来分配内存

// new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针
// make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身

package main

import "fmt"

func main() {
	p := fmt.Println
	fmt.Println(p)
	var ptr *func(...interface{}) (int, error) = &p
	fmt.Println(ptr, *ptr)
	*ptr = fmt.Print
	fmt.Println(ptr, *ptr)

	var a *int = new(int)
	var b *int = new(int)
	fmt.Println(a, *a, b, *b)
	*a = 100
	fmt.Println(*a)

	//var b map[string]int
	var d = map[string]int{"1": 1}
	//b := map[string]int{"1": 1}
	d["测试"] = 100
	fmt.Println(d)
}
