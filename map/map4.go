package main

import "fmt"

type Test struct {
	Name string
}

var list map[string]Test

func main() {

	list = make(map[string]Test)
	name := Test{"xiaoming"}
	//name.Name = "he"
	list["name"] = name
	list["name"].Name = "Hello"
	fmt.Println(list["name"])
}

// 因为list[“name”]不是一个普通的指针值，
// map的value本身是不可寻址的，因为map中的值会在内存中移动，并且旧的指针地址在map改变时会变得无效。

// 如果实在需要类似的功能，可改为var list map[string]*Test，参见map5_4
