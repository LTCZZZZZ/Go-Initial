package main

import (
	"MyTools"
	"fmt"
)

func main()  {
	is1 := MyTools.Pyrange(1, 10)
	is2 := MyTools.Pyrange(1, 10)
	is3 := MyTools.Pyrange(1, 5)
	iss := [][]int{is1, is2}
	fmt.Println(iss)
	for _, is := range iss {
		is[0] = 100
		is = is3  // 引用传递。另外，此时is已经不是iss中的is了
		is[0] = 100
	}
	fmt.Println(iss)
	fmt.Println(is3)
}
