package main

import "fmt"

func main() {
	var (
		a uint8 = 15
		b uint8 = 50
	)
	fmt.Println(a * b) // 238，因为超出运算范围了

	var (
		c uint32 = 15
		d uint32 = 50
	)
	fmt.Println(c * d) // 750
}
