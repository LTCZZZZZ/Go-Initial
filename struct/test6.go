// 任意类型添加方法
// 注意事项： 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。
// 比如，不能直接为内置类型int添加方法，
// 但对于强类型语言如go来说，用MyInt替代int并自定义额外的方法，似乎没有太大影响，因为变量声明时本来就需要声明类型

package main

import "fmt"

//MyInt 将int定义为自定义MyInt类型
type MyInt int

//SayHello 为MyInt添加一个SayHello的方法
func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个int。")
}
func main() {
	var m1 MyInt
	m1.SayHello() //Hello, 我是一个int。
	m1 = 100
	fmt.Printf("%#v  %T\n", m1, m1) //100  main.MyInt
}
