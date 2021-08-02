package main

import "fmt"

func main() {
	s := "\n \r\n \t"
	b := []byte(s)
	//for _, i := range b {
	//	fmt.Println(i == 10)
	//}
	fmt.Println(b)
	fmt.Println(len(s))

}
