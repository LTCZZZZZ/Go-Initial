// Python递归有默认深度限制，Golang是否有递归深度限制？
package main

func Test1() {
	//fmt.Println(time.Now(), "Test1")
	//time.Sleep(500 * time.Millisecond)
	Test1() // runtime: goroutine stack exceeds 1000000000-byte limit
}

func main() {
	Test1()
}
