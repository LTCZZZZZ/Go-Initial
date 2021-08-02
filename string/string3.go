package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str1 := "ABC"
	fmt.Println(utf8.ValidString(str1)) // true

	str2 := "A\xfeC"
	fmt.Println(utf8.ValidString(str2)) // false
	fmt.Println(str2)

	str3 := "A\\xfeC cafe\u0301 café"
	fmt.Println(utf8.ValidString(str3)) // true    // 把转义字符转义成字面值
	fmt.Println(str3)

	str4 := "é"
	fmt.Println(len(str4))                    // 2
	fmt.Println(len([]rune(str4)))            // 1
	fmt.Println(utf8.RuneCountInString(str4)) // 1，某些文档中会说结果是2，但实测结果是1

}
