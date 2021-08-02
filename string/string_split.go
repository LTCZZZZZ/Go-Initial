package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// 可以看出，中文的byte流第一个数字大都位于228-231之间，
	// 可以猜测，在byte流中，某一范围开头的数字表示它是某种文字，然后再按这种文字的规则向后取数据去解析
	s := "[]#我的中华0\na\nb\nc"
	b := []byte(s)
	//for _, i := range b {
	//	fmt.Println(i == 10)
	//}
	fmt.Println(b)
	fmt.Println(len(s))
	b[0] = 92
	fmt.Println(string(b))
	fmt.Println(strings.Split(s, "\n"))
	fmt.Println(bytes.Split(b, []byte("\n")))

	fmt.Println(strings.Trim(s, "#"))

}
