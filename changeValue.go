package main

import (
	"fmt"
)

/* 定义结构体 */
type Circle struct {
	radius float64
}

func main() {
	var c Circle
	//c := &Circle{}
	fmt.Println(c)
	fmt.Println(c.radius)
	c.radius = 10.00
	fmt.Println(c.getArea())

	// 在c直接定义为Circle时，以下两种方式所得结果相同，而如果c := &Circle{}，第二行就用不了了
	c.changeRadius(20)
	//(&c).changeRadius(20)

	fmt.Println(c.radius)
	change(&c, 30)
	fmt.Println(c.radius)
}
func (c Circle) getArea() float64 {
	return c.radius * c.radius
}

// 注意如果想要更改成功c的值，这里需要传指针
func (c *Circle) changeRadius(radius float64) {
	// 重要：之所以下面这一行代码对不同定义的c都生效，是因为：
	// 可以使用结构体指针访问结构体成员，只要使用 "." 操作符
	// 此外，貌似对结构体指针定义的方法，结构体本身也能直接调用（虽然使用起来方便了，但这看起来会显得有些混乱）
	c.radius = radius
}

// 以下操作将不生效
//func (c Circle) changeRadius(radius float64)  {
//  c.radius = radius
//}

// 引用类型要想改变值需要传指针
func change(c *Circle, radius float64) {
	c.radius = radius
}

// 另一个例子
type Rect struct { //定义矩形类
	x, y          float64 //类型只包含属性，并没有方法
	width, height float64
}

func (r *Rect) Area() float64 { //为Rect类型绑定Area的方法，*Rect为指针引用可以修改传入参数的值
	return r.width * r.height //方法归属于类型，不归属于具体的对象，声明该类型的对象即可调用该类型的方法
}
