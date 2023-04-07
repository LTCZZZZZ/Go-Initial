package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

//Class 班级
type Class struct {
	Title    string
	Students []*Student
}

func main() {
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 3; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)

	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"}]}`

	// 必须得赋一个初值或者new一个，否则会报错，因为c1是一个指针，如果不赋值，那么c1就是nil pointer
	//c1 := &Class{}
	var c1 *Class
	c1 = new(Class)
	fmt.Println(c1)

	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println(err)
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)

	fmt.Println(`
How are you?
Fine, thank you. And you?
I'm fine too.
        `)
	fmt.Println(`end`)

	var s interface{}
	err = json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(s)
}
