// 闭包 closure
package test

import (
	"fmt"
	"testing"
)

func getSequence() func() int {
	i := 0
	getNumber := func() int {
		i += 1
		return i
	}
	fmt.Println(i)  // 0
	return getNumber
}

func TestGetSequence(t *testing.T) {
	g := getSequence()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	g2 := getSequence()
	fmt.Println(g2())
	fmt.Println(g2())
}
