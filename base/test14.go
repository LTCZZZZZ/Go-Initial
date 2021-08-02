package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

//type H map[string]interface{}
type H map[interface{}]interface{} // 这样定义的map就和python中的dict没什么区别了
// 个人猜测，interface{}相当于一个空接口(因为里面没有实现任何函数)，所以它可以代表任何类型
type Empty interface{}

//可以定义一个空接口类型的array、slice、map、struct等，这样它们就可以用来存放任意类型的对象

func main() {

	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})

	d := map[int]int{}
	d[1] = 2
	d[3] = 5
	fmt.Println(d)

	d2 := H{"C": 3, "5": H{"inner": 2}, "6": d, 6: 3}
	fmt.Println(d2)

	e := new(Empty)
	fmt.Println(e)

	// 下面两种定义效果相同
	//var e1, e2, e3 Empty
	var e1, e2, e3 interface{}
	e1 = 3
	e2 = 'z'
	e3 = "字符串"
	fmt.Println(e1, e2, e3)
	e1 = "int 3变成 字符串"
	e2 = 99
	e3 = 12345
	fmt.Println(e1, e2, e3)
}

func test1(int, int) int {
	return 0
	//return aa + bb

}
