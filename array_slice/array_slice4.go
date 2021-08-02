// 超出原 slice.cap 限制，就会重新分配底层数组，即便原数组并未填满
package main

import (
	"fmt"
)

func main() {

	data := [...]int{0, 1, 2, 3, 4, 10: 0}
	s := data[:2:3]              // 此处如果不限制s的容量，则最后打印指针的地址会相同
	fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针。相同
	fmt.Println(cap(s))

	s = append(s, 100, 200) // 一次 append 两个值，超出 s.cap 限制。

	fmt.Println(s, data)         // 重新分配底层数组，与原数组无关。
	fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针。不同

}
