//除用 panic 引发中断性错误外，还可返回 error 类型错误对象来表示函数调用状态。

package main

import (
	"errors"
	"fmt"
)

var ErrDivByZero = errors.New("division by zero")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivByZero
	}
	return x / y, nil
}

func main() {
	defer func() {
		fmt.Println(recover())
	}()
	switch z, err := div(10, 0); err {
	case nil:
		println(z)
	case ErrDivByZero:
		panic(err) // 假如此处不用panic而改用下面的打印的话，上面的defer recover()结构就不需要
		//fmt.Println(err)
	}
}
