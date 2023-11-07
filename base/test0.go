package main

import (
	"fmt"
	"strings"
)

type Mini struct {
	ID int
}

func main() {
	var position [][]float64
	fmt.Println(position)
	fmt.Println(position == nil)
	fmt.Println(len(position))
	//position = position[0:0]
	position = make([][]float64, 0)
	fmt.Println(position == nil)

	position = [][]float64{{0, 1}, {2, 3}}
	position = position[0:0]
	fmt.Println(position == nil)

	var (
		f float64
	)
	f = 200.0
	fmt.Println(f)
	fmt.Println(fmt.Sprintf("%v", f)) // 正常打印及%v不会有小数点
	fmt.Println(f == 200)             // 这里200是被当做了浮点数，所以是true

	// 对定义好的结构，哪怕获取不到，数据也会有默认值0，不会报错
	var miniMap map[string]Mini
	m := miniMap["asd"]
	fmt.Println(m.ID)

	fmt.Println("8.0.28" >= "8.0")
	fmt.Println("5.7" >= "8.0")

	var mm map[string]interface{}
	fmt.Println(mm == nil)

	// 可以看到，一般都是值传递，但是如果是slice，就是引用传递
	var s string = "origin"
	changeParameterString(s)
	fmt.Println(s)

	var i int = 1
	changeParameterInt(i)
	fmt.Println(i)

	var sl []int = []int{1, 2, 3}
	changeParameterSlice(sl)
	fmt.Println(sl)

	z := strings.SplitN("abc", "-", 2)
	fmt.Println(z)

	var err error = nil
	fmt.Println(err == nil)
	// err为nil时，err.Error()会报错
	//fmt.Println(err.Error())
	//err = fmt.Errorf("this is error")
	fmt.Println(fmt.Sprintf("%v", err))

}

func changeParameterString(s string) {
	s += "changed"
	fmt.Println(s)
}

func changeParameterInt(i int) {
	i = 100
}

func changeParameterSlice(s []int) {
	s = []int{4, 5, 6} // 这里不会改变原来的值，因为是将变量指向了新的地址
	s[0] = 100         // 这里会改变原来的值，但如果有上一行，外层s的值不会改变，因为上一行已经将s指向了新的地址
}
