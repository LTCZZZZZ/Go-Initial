package main

import "fmt"

type Duration int64

const (
	minDuration Duration = -1 << 63
	maxDuration Duration = 1<<63 - 1
)

func main()  {
	fmt.Println(minDuration)
	fmt.Println(maxDuration)
	var res []int
	res = append(res, 1, 2, 3)
	fmt.Println(res)
}
