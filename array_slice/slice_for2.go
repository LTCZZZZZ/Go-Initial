package main

import (
	"MyTools"
	"fmt"
)

func main() {
	is1 := MyTools.Pyrange(1, 3)
	is2 := MyTools.Pyrange(1, 5)
	//is3 := MyTools.Pyrange(1, 7)
	iss := [][]int{is1, is2}
	fmt.Println(iss)
	for j := 0; j < 2; j++ {
		for _, is := range iss {
			//is = MyTools.Pyrange(1, 10)  // 如果有了这一行，下面的is=append操作对iss也无效了
			for i := range is {
				if i == 0 {
					// 细想想下面这一行代码为什么能影响到iss，它又做了些什么
					// 首先，在range中is是值传递，直接对is赋值的话是无效的，
					// 但is指向底层数组，下面的操作修改了底层数组，
					// 下面三种方式对iss的影响相同
					//is = append(is[:i], is[i+1:]...)
					// 甚至可以改成如下两种形式
					//b := append(is[:i], is[i+1:]...)
					//fmt.Println(b)
					// 或者更直接一点
					_ = append(is[:i], is[i+1:]...)  // 这种方式看起来最顺眼
				}
			}
		}
		fmt.Println(iss)
	}
	//fmt.Println(iss)
	//fmt.Println(is3)

	fmt.Println(iss)

}
