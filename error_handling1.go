package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	a := 1.0
	//b := 0.8
	//fmt.Println(Sqrt(-a))
	res, err := Sqrt(a - 3)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func Sqrt(f float64) (res float64, err error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	res = math.Sqrt(f)
	return
}
