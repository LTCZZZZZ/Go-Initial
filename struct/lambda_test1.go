package main

import "fmt"

//Address 地址结构体
type Address struct {
	Province   string
	City       string
	CreateTime string
}

//User 用户结构体
type User struct {
	Name    string
	Gender  string
	Address //匿名结构体
	Email
}

//Email 邮箱结构体
type Email struct {
	Account    string
	CreateTime string
}

func main() {
	var user2 User
	user2.Name = "pprof"
	user2.Gender = "女"
	user2.Address.Province = "黑龙江"   //通过匿名结构体.字段名访问
	user2.City = "哈尔滨"               //直接访问匿名结构体的字段名
	fmt.Printf("user2=%#v\n", user2) //user2=main.User{Name:"pprof", Gender:"女", Address:main.Address{Province:"黑龙江", City:"哈尔滨"}}

	var user3 User
	user3.Name = "pprof"
	user3.Gender = "女"
	//user3.CreateTime = "2019" //ambiguous selector user3.CreateTime
	user3.Address.CreateTime = "2000" //指定Address结构体中的CreateTime
	user3.Email.CreateTime = "2000"   //指定Email结构体中的CreateTime
}

// 当访问结构体成员时会先在结构体中查找该字段，找不到再去匿名结构体中查找。这种特性使结构体可以方便的继承

// 嵌套结构体内部可能存在相同的字段名。这个时候为了避免歧义需要指定具体的内嵌结构体的字段。
