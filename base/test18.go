package main

import (
	"fmt"
	"strconv"
)

func main() {
	val, err := strconv.Atoi("123f")
	fmt.Println(val, err)
}
