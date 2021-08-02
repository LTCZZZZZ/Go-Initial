// 重要
// 当要将结构体对象转换为 JSON 时，对象中的属性首字母必须是大写，才能正常转换为 JSON，
// 首字母小写的属性不会包含在转为json后的数据内(这也是可以善加利用的地方之一)

// 那这样 JSON 字符串以后就只能是大写了么？ 当然不是，可以使用 tag 标记要返回的字段名。
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name string //Name字段首字母大写
	age  int    //age字段首字母小写
}

type Person2 struct {
	Name string //都是大写
	Age  int
}

type Person3 struct {
	Name string `json:"name"` //标记json名字为name
	Age  int    `json:"age"`
	Time int64  `json:"-"` // 标记忽略该字段

}

func main() {
	person := Person{"小明", 18}
	if result, err := json.Marshal(&person); err == nil { //json.Marshal 将对象转换为json字符串
		fmt.Println(string(result))
	}
	person2 := Person2{"小明", 18}
	if result, err := json.Marshal(&person2); err == nil { //json.Marshal 将对象转换为json字符串
		fmt.Println(string(result))
	}
	person3 := Person3{"小明", 18, time.Now().Unix()}
	if result, err := json.Marshal(&person3); err == nil { //json.Marshal 将对象转换为json字符串
		fmt.Println(string(result))
	}
}
