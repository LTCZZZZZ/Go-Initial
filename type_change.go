package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)

	i := int(mean)
	f := 3.8
	s1 := strconv.FormatFloat(f, 'f', -1, 64)
	s2 := strconv.FormatFloat(f, 'f', 5, 64)
	s3 := strconv.FormatFloat(f, 'g', 5, 64)
	fmt.Printf("%T\n", f)
	fmt.Println(i, int(f))
	fmt.Println(s1, s2, s3)
}
