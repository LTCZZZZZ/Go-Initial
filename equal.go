package main

import (
	"fmt"
	"reflect"
)

type User struct {
	AddTime int
	Name    string
	Age     int
}

func main() {
	fmt.Println(reflect.ValueOf(3) == reflect.ValueOf(2+1))                                       // true
	fmt.Println(reflect.DeepEqual(3, 2+1))                                                        // true
	fmt.Println(reflect.DeepEqual(reflect.ValueOf(3), reflect.ValueOf(2+1)))                      // true
	fmt.Println(3 == User{3, "三", 30}.AddTime)                                                    // true
	fmt.Println(reflect.DeepEqual(3, User{3, "三", 30}.AddTime))                                   // true
	fmt.Println(reflect.ValueOf(3) == reflect.ValueOf(User{3, "三", 30}.AddTime))                  // false
	fmt.Println(reflect.DeepEqual(reflect.ValueOf(3), reflect.ValueOf(User{3, "三", 30}.AddTime))) // false
	fmt.Println(3 == reflect.ValueOf(User{3, "三", 30}.AddTime).Int())                             // true
	fmt.Println(reflect.DeepEqual(3, reflect.ValueOf(User{3, "三", 30}.AddTime).Int()))            // false，因为后者返回int64
	fmt.Println(reflect.DeepEqual(int64(3), reflect.ValueOf(User{3, "三", 30}.AddTime).Int()))     // true
	fmt.Println(3 == reflect.ValueOf(User{3, "三", 30}.AddTime).Interface())                       // true
	fmt.Println(reflect.DeepEqual(3, reflect.ValueOf(User{3, "三", 30}.AddTime).Interface()))      // true

	fmt.Println(3 == int64(3))                  // true
	fmt.Println(3 == int32(3))                  // true
	fmt.Println(reflect.DeepEqual(3, int64(3))) // false
	fmt.Println(reflect.DeepEqual(3, int32(3))) // false
	fmt.Println(reflect.DeepEqual(3, int(3)))   // true
}
