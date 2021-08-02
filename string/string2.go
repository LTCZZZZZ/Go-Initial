package main

import (
	"MyTools"
	"fmt"
	"unicode/utf8"
)

func traversalString() {
	s := "pprof.cn博客en"
	fmt.Println("length:", len(s))
	fmt.Println("chinese word length:", len([]rune(s)))
	fmt.Println("chinese word length:", utf8.RuneCountInString(s))

	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v:%v(%c) ", i, s[i], s[i])
	}
	fmt.Println()

	for i, r := range s { //rune
		fmt.Printf("%v:%v(%c) ", i, r, r)
	}
	fmt.Println()

	for i, r := range []rune(s) { //rune，但索引和上面的不同，（需特别注意）
		fmt.Printf("%v:%v(%c) ", i, r, r)
	}
	fmt.Println()

	for i, r := range []int32(s) { //int32，结果和rune一模一样，int32每个都占4字节，故而，转成列表时，是先切分，后转的
		if i == 0 {
			r = 23849
		}
		fmt.Printf("%v:%v(%c) ", i, r, r)

	}
	fmt.Println()
	s_i := []int32(s)
	s_i[0] = 23849
	fmt.Println(string(s_i))
	fmt.Println(string([]int32{23442, 49232, 28373, 23985}))
	fmt.Println(string(MyTools.PyrangeInt32(21330, 21350)))
}

func changeString() {
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))
}

func main() {
	traversalString()
	changeString()
}
