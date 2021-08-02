package main

import (
	"fmt"
	"reflect"
)

type Dog struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	// 下面两种都可以
	//var d interface{}
	//d = &Dog{}
	//fmt.Println(change(d))

	var d = Dog{}
	fmt.Println(change(&d))

	fmt.Println(d)
}

func change(v interface{}) (interface{}, interface{}) {

	rDog := reflect.ValueOf(v).Elem()
	vAge := rDog.FieldByName("Age")
	vAge.SetInt(1)

	vSex := rDog.FieldByName("Sex")
	vSex.SetString("male")
	//res1 := reflect.ValueOf(v).Elem().Interface()
	res1 := rDog.Interface() // 获取指针变量内的值到interface类型
	vAge.Set(reflect.ValueOf(2))
	//res2 := reflect.ValueOf(v).Elem().Interface()
	res2 := rDog.Interface()
	return res1, res2
}
