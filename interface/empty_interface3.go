package main

import "fmt"

func main() {
	m := make(map[string]interface{})
	m["a"] = "AA"
	fmt.Println(m)
	fmt.Println(m["b"] == nil)
	v, ok := m["b"]
	fmt.Println(v, ok)
}
