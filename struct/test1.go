package main

import "fmt"

type H map[string]interface{} // 类型定义
//type H = map[string]interface{}  // 类型别名

type person struct { // 结构体
	name string
	city string
	age  int8
	H    H // 第一个H可省略，即 匿名结构体
}

func main() {
	a := H{"test": 12}
	fmt.Printf("%T\n", a)
	fmt.Println(a)
	a = H{"test": a}
	fmt.Printf("%T\n", a)
	fmt.Println(a)

	p := person{"zs", "bj", 50, a}
	fmt.Printf("%T\n", p)

	//var p2 person  // 初始化会用默认值填充
	//fmt.Println(p2)

	var p2 = new(person)
	//p2 := &person{}  // 和上面的等价
	fmt.Printf("%T\n", p2)     //*main.person
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}

	// Go语言中支持对结构体指针直接使用.来访问结构体的成员
	p2.name = "测试"
	p2.age = 18
	p2.city = "北京"
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"测试", city:"北京", age:18}
}
