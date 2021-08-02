package main

import (
	"fmt"
)

func main() {
	var s1 = "hello world"
	var b = []byte(s1) // 字符串转字节切片
	var s2 = string(b) // 字节切片转字符串
	fmt.Println(b)
	fmt.Println(s2)
	fmt.Println(s1 == s2)
	fmt.Println(len(s1))

	s1 = "嘻哈china"
	fmt.Println(s1[3:6])

	fmt.Println(len(s1))

	j := 0
	for i, s := range s1 {
		//if i == 0 {
		//	continue
		//}
		fmt.Println(i, s, s1[i], fmt.Sprintf("%x", s1[i]), s1[j:i])
		j = i
	}

	fmt.Println("____________")
	for i := 0; i < len(s1); i++ {
		fmt.Println(i, s1[i], fmt.Sprintf("%x", s1[i]))
	}
}
