package main

import (
	"fmt"
	"sort"
)

// 对结构体指定字段进行排序
type User struct {
	Name string `json:"name"` // `json:"xxx"`：在结构体和json字符串字段顺序不一致的情况下：unmarshal根据tag去寻找对应字段的内容
	Age  int    `json:"age"`
}

func main() {
	users := []User{
		{
			Name: "test1",
			Age:  22,
		},
		{
			Name: "test2",
			Age:  19,
		},
		{
			Name: "test3",
			Age:  25,
		},
	}

	sort.Slice(users, func(i, j int) bool { // desc
		return users[i].Age > users[j].Age
	})
	fmt.Printf("按Age降序：%+v\n", users)

	sort.SliceStable(users, func(i, j int) bool { // asc
		return users[i].Age < users[j].Age
	})
	fmt.Printf("按Age升序：%+v\n", users)

	fmt.Println(users[0:0])
}

// 此外也可使用sort.Sort()方法，不过需要自己去实现 Len()、Swap()、Less()方法，
// 参考：golang对自定义类型排序 https://segmentfault.com/a/1190000008062661