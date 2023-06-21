package main

import "fmt"

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
	fmt.Println(f == 200) // 这里200是被当做了浮点数，所以是true

	// 对定义好的结构，哪怕获取不到，数据也会有默认值0，不会报错
	var miniMap map[string]Mini
	m := miniMap["asd"]
	fmt.Println(m.ID)

}
