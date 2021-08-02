package main

import (
	"fmt"
)

func main()  {
	var a []int
	fmt.Printf("%p\n", a)
	a = []int{1, 2, 3, 4, 5}
	fmt.Printf("%p %p %p %p\n", a, a[1:], a[2:], a[3:])  // 每个相差8位
	//a = a[1:] // 删除开头1个元素
	a = a[:copy(a, a[1:])]
	fmt.Println(a)
	fmt.Printf("%p\n", a)  // 直接是a[1:]的地址

	//a = append(a[:i], a[i+1:]...) // 删除中间1个元素
	//a = append(a[:i], a[i+N:]...) // 删除中间N个元素
	//a = a[:i+copy(a[i:], a[i+1:])] // 删除中间1个元素
	//a = a[:i+copy(a[i:], a[i+N:])] // 删除中间N个元素

	a = []int{1, 2, 3, 4, 5}
	fmt.Printf("%p %p %p %p\n", a, a[1:], a[2:], a[3:])  // 每个相差8位
	b := append(a[:1], a[2:]...)
	fmt.Printf("%p %p %p %p\n", a, a[1:], a[2:], a[3:])  // 地址没变，相当于移位了
	fmt.Println(b)
	fmt.Println(a)
	b[0] = 5
	fmt.Println(a)

}
