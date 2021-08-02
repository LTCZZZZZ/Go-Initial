package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	var t interface{}

	t = &T{18, "nick"}
	fmt.Println()
	s := reflect.ValueOf(t)
	fmt.Println(s)
	typeOfT := s.Type()
	fmt.Println(typeOfT)
	//fmt.Println(s.NumField())

	//for i := 0; i < s.NumField(); i++ {
	//	f := s.Field(i)
	//	fmt.Printf("%d: %s %s = %v\n", i,
	//		typeOfT.Field(i).Name, f.Type(), f.Interface())
	//}

	//s.Field(0).SetInt(25)
	//s.Field(1).SetString("nicky")
	reflect.ValueOf(t).Elem().FieldByName("B").Set(reflect.ValueOf("wode"))
	//reflect.ValueOf(&t).Elem().FieldByName("B").Addr().Interface()
	fmt.Println(t)
}
