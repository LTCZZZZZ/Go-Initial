// panic
// 1、内置函数
// 2、假如函数g中书写了panic语句，会终止其后要执行的代码，在panic所在函数g内如果存在要执行的defer函数列表，按照defer的逆序执行
// 3、返回函数g的调用者f，在f中，调用函数f语句之后的代码不会执行，假如函数f中存在要执行的defer函数列表，按照defer的逆序执行，
//    但如果g中defer定义了recover()，则g最后正常退出，f正常执行
// 4、直到goroutine整个退出，并报告错误

// recover
// 1、内置函数
// 2、用来控制一个goroutine的panicking行为，捕获panic，从而影响应用的行为
// 3、一般的调用建议
//     a). 在defer函数中，通过recever来终止一个goroutine的panicking过程，从而恢复正常代码的执行
//     b). 可以获取通过panic传递的error

// 注意
// 1.利用recover处理panic指令，defer 必须放在 panic 之前定义，另外 recover 只有在 defer 调用的函数中才有效。否则当panic时，recover无法捕获到panic，无法防止panic扩散。
// 2.recover 处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点。
// 3.多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调用。

package main

import "fmt"

func f() {
	defer fmt.Println("这段会执行，在f退出之前")
	// defer函数放到main中也能生效，只要在f执行之前
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
		//fmt.Println("不处理错误！")  // 如果注释掉上面这一段改为这一段，这会执行这一段，执行完后报错
	}()
	g()
	fmt.Println("这段代码也不会执行，除非g中有defer recover()")
}

func g() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in g", r)
		}
		//fmt.Println("不处理错误！")  // 如果注释掉上面这一段改为这一段，这会执行这一段，执行完后报错
	}()
	a := 1
	b := 0
	//panic("panic error!")  // 也可以主动抛出错误
	fmt.Println(a / b) // 不能直接写1 / 0，否则编译时就通不过
	fmt.Println("这段代码不会执行")
}

func main() {
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("Recovered in f", r)
	//	}
	//}()
	f()
	fmt.Println("这段代码会执行，因为f是正常运行退出的")
}
