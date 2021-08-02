// 结构体相关反射机制研究，reflect
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func (u *User) SetAge(a int) {
	u.Age = a
}

func main() {
	u := User{"天天", 0}
	n := "Name"
	fmt.Println(reflect.ValueOf(u).FieldByName(n))
	fmt.Println(reflect.ValueOf(u).Field(1))
	//fmt.Println(reflect.ValueOf(u).Method(0))
	fmt.Println(reflect.ValueOf(&u).Method(0)) // 由于上面方法定义的是*，所以这里必须要&，下面也一样
	reflect.ValueOf(&u).MethodByName("SetAge").Call([]reflect.Value{reflect.ValueOf(10)})
	fmt.Println(u)

	var v interface{}
	v = User{"天天", 0}
	//value := v.(User)
	fmt.Println(reflect.ValueOf(v).FieldByName(n))
	//reflect.ValueOf(&v).MethodByName("SetAge").Call([]reflect.Value{reflect.ValueOf(20)})  // 这样不行

	// 想在不知道其类型的时候进行类型装换，做不到
	//value := v.(type)，因为value不知道类型就无法初始化，即获得变量的地址和内存分配。只能显示转换

	//v = &v  // 就成了*interface {}
	v = &User{"天天", 0}
	fmt.Println(reflect.TypeOf(v))
	reflect.ValueOf(v).MethodByName("SetAge").Call([]reflect.Value{reflect.ValueOf(20)})

	//value := v.(reflect.TypeOf(v))
	//fmt.Println(value)
	fmt.Println(v)
	fmt.Println(v.(*User))

	v = []interface{}{1, 2, 3}
	fmt.Println(v.([]interface{})[1])

	var a, b interface{}
	a = 1
	b = uint(1)
	fmt.Println(a == b) // 空接口中类型相同，值相同，则相等

	//fmt.Println(reflect.ValueOf(a).(interface{})) // Value类型不能进行类型转换了
	fmt.Println(reflect.ValueOf(a).Interface())

	//fmt.Println(int(3) == int64(3))  // 编译就过不了

	a = int(3)
	c := uint(3)
	fmt.Println(uint32(c) == a)
	fmt.Println(int(c) == a)
}
