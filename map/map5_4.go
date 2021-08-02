package main

import "fmt"

type Test2 struct {
	Name string
}

var list2 map[string]*Test2

func main() {

	list2 = make(map[string]*Test2)
	name := Test2{"xiaoming"}
	//name.Name = "he"
	list2["name"] = &name
	list2["name"].Name = "Hello"
	fmt.Println(list2["name"])
}

// 因为list[“name”]不是一个普通的指针值，
// map的value本身是不可寻址的，因为map中的值会在内存中移动，并且旧的指针地址在map改变时会变得无效。

// 如果实在需要类似的功能，将map的value直接设为指针类型
