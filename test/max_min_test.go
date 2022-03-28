package test

import (
	"MyTools"
	"fmt"
	"testing"
	"time"
)

func TestMaxMin(t *testing.T) {
	//a := MyTools.Pyrange(1000000, 0, -1)  //0.061106 s
	//a := MyTools.Pyrange(1000000)  // 0.061451s，比python快了约100倍，唉
	a := MyTools.PyrangeInt32(1000000)  // 改为int32还能更快，0.057557s
	//fmt.Println(a[0: 3], a[100000], len(a))

	start_time := time.Now().UnixNano()
	fmt.Println(start_time)

	for i := 0; i < 100; i++ {
		// 待测试的代码
		var big int32 = 0
		for _ ,i := range a {  // 这里也是i，但变量的作用域不同，在这里用i的话，循环内就获取不到外层i的值了，不过此处无影响
			if i > big {
				big = i
			}
		}
		fmt.Println(big)
	}

	end_time := time.Now().UnixNano()
	duration := end_time - start_time
	fmt.Println(duration)
	fmt.Println(float64(duration) / 1000000000)
}
