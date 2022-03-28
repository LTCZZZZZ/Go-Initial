package test

import (
	"fmt"
	"testing"
)

type Duration int64

const (
	minDuration Duration = -1 << 63
	maxDuration Duration = 1<<63 - 1
)

func Test(t *testing.T) {
	fmt.Println(minDuration)
	fmt.Println(maxDuration)
	var res []int
	res = append(res, 1, 2, 3)
	fmt.Println(res)
	var a interface{}
	a = int64(3)
	fmt.Println(uint32(3) == a)
	a = 3
	fmt.Println(uint32(3) == a)
	a = uint32(3)
	fmt.Println(uint32(3) == a)
}
