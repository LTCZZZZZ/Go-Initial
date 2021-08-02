package main

import "fmt"

//Person 结构体
type Person struct {
	name string
	age  int8
}

//NewPerson 构造函数，不知道这个有什么用，查了一下似乎可用于添加默认值之类的，
// 但是那样添加一个特定的初始化方法感觉会更好
// 下面这个函数也可不返回指针类型，貌似不影响
// 但是，因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//Dream Person做梦的方法
// 方法与函数的区别是，函数不属于任何类型，方法属于特定的类型
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

// SetAge 设置p的年龄
// 使用指针接收者
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

// SetAgeFail 设置p的年龄，对函数外无效
// 不使用指针接收者，修改对外层对象无效
// 当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。
//在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。
func (p Person) SetAgeFail(newAge int8) {
	p.age = newAge
}

func main() {
	p1 := NewPerson("测试", 25)
	fmt.Printf("%T\n", p1)
	p1.Dream()
	p1.SetAge(30)
	fmt.Println(p1.age) // 30
	p1.SetAgeFail(40)
	fmt.Println(p1.age) // 30
}

// 可用实例 value 或 pointer 调用全部方法，编译器自动转换。
// 用<type>.func(instance, args...)调用时可能受到方法集的限制，参见 http://www.topgoer.com/%E6%96%B9%E6%B3%95/%E6%96%B9%E6%B3%95%E9%9B%86.html
// 但，对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然。
// 注意，当接受者不是一个指针时，该方法操作对应接受者的值的副本
//(意思就是即使你使用了指针调用函数，但是函数的接受者是值类型，所以函数内部操作还是对副本的操作，而不是指针操作。

// 另外，不是指针时调用方法会复制原结构体传入函数，如果结构体很大，会有较大的性能开销

// 什么时候应该使用指针类型接收者?
// 1.需要修改接收者中的值
// 2.接收者是拷贝代价比较大的大对象
// 3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
