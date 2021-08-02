package main

import "fmt"

func main() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}

	println()

	//从以下代码输出的结果可以看出：
	//switch 从第一个判断表达式为 true 的 case 开始执行，
	//如果 case 带有 fallthrough，程序会继续执行下一条 case，
	//且它不会去判断下一个 case 的表达式是否为 true。

	//switch {
	//case false:
	//	fmt.Println("1、case 条件语句为 false")
	//	fallthrough
	//case true:
	//	fmt.Println("2、case 条件语句为 true")
	//	fallthrough
	//case false:
	//	fmt.Println("3、case 条件语句为 false")
	//	fallthrough
	//case true:
	//	fmt.Println("4、case 条件语句为 true")
	//case false:
	//	fmt.Println("5、case 条件语句为 false")
	//	fallthrough
	//default:
	//	fmt.Println("6、默认 case")
	//}
}
