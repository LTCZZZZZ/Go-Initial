package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//Student 学生
type Student2 struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

func main() {
	s1 := Student2{
		ID:     1,
		Gender: "女",
		name:   "pprof",
	}
	//data, err := json.Marshal(s1)
	data, err := json.Marshal(reflect.ValueOf(s1).Interface())
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"女"}
}
