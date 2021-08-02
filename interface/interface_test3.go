package main

import "fmt"

type fruit interface {
	getName() string
	setName(name string)
}
type apple struct {
	name string
}

//[1]
func (a apple) getName() string {
	return a.name
}

//[2]
func (a *apple) setName(name string) {
	a.name = name
}

type pear struct {
	name string
}

func (a pear) getName() string {
	println(a.name)
	return a.name
}

func (a pear) setName(name string) {
	a.name = name // 这里对外部是不生效的
	println(a.name)
}

func main() {
	var a fruit
	a = &apple{"红富士"} // 注意此处必须用&，否则会报错，感觉不是apple类而是(&apple)类implement了fruit接口
	fmt.Printf("%T\n", a)
	fmt.Println(a.getName())
	a.setName("树顶红")
	fmt.Println(a.getName())

	a = pear{"库尔勒"} // 而此处不用&
	a.setName("新库尔勒")
	a.getName()
}
