package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	fmt.Println(1)
	//debug.PrintStack()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println(string(debug.Stack()))
		}
	}()

	panic("manual panic error")

	fmt.Println(2)
}
