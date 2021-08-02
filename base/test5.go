//指针相关
package main

import "fmt"

//func main() {
//	var a int = 4
//	var b int32
//	var c float32
//	var ptr *int
//
//	/* 运算符实例 */
//	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a );
//	fmt.Println("第 1 行 - a 变量类型为 = %T\n", a, 100, 200);
//	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b );
//	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c );
//
//	/*  & 和 * 运算符实例 */
//	ptr = &a     /* 'ptr' 包含了 'a' 变量的地址 */
//	println(ptr)
//	fmt.Println(ptr)
//	fmt.Printf("a 的值为  %d\n", a);
//	fmt.Printf("*ptr 为 %d\n", *ptr);
//
//}

// 使用指针变量与不使用的区别
func main() {
	var a int = 4
	var ptr int
	ptr = a
	fmt.Println(ptr) //4
	a = 15
	fmt.Println(ptr) //4

	var b int = 5
	var ptr1 *int
	ptr1 = &b
	fmt.Println(*ptr1) //5
	b = 15
	fmt.Println(*ptr1) //15
}
