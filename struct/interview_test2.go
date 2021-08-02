package main

import "fmt"

type student struct {
	id   int
	name string
	age  int
}

func demo(ce []student) {
	//切片是引用传递，是可以改变值的
	ce[1].age = 999
	// ce = append(ce, student{3, "xiaowang", 56})
	// return ce
}
func main() {
	var ce []student //定义一个切片类型的结构体
	ce = []student{
		student{1, "xiaoming", 22},
		student{2, "xiaozhang", 33},
	}
	fmt.Println(ce)
	demo(ce)
	fmt.Println(ce)
	fmt.Printf("%v\n", ce)

	ce2 := make(map[int]student)
	ce2[1] = student{1, "xiaolizi", 22}
	ce2[2] = student{2, "wang", 23}
	fmt.Println(ce2)
	delete(ce2, 2)
	fmt.Println(ce2)
}
