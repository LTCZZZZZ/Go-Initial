package main

import "fmt"

func main()  {
	var position [][]float64
	fmt.Println(position)
	fmt.Println(position == nil)
	fmt.Println(len(position))
	//position = position[0:0]
	position = make([][]float64, 0)
	fmt.Println(position == nil)

	position = [][]float64{{0,1},{2,3}}
	position = position[0:0]
	fmt.Println(position == nil)
}
