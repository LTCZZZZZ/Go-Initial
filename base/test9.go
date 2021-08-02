// For-each range 循环
// 这种格式的循环可以对字符串、数组、切片等进行迭代输出元素
package main

import "fmt"

func main() {
	strings := [4]string{"google", "runoob"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 5}
	numbers[5] = 10
	println(numbers[3])
	for i, x := range numbers {
		// Goland在查看时会自动转换为合理的显示方式
		//fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

}
