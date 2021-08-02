package main

import "fmt"

func main() {
	a := make(map[string]interface{})
	a["1"] = 1
	a["2"] = 2
	b := a
	b["1"] = 5
	b["3"] = 3
	fmt.Println(a)
	fmt.Println(b)

}
