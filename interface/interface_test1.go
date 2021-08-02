package main

//
//import (
//	"fmt"
//)
//
//type Phone interface {
//	call(year int) string
//	callSelf() string
//}
//
//type NokiaPhone struct {
//	Year int
//}
//
//func (nokiaPhone NokiaPhone) call(year int) string {
//	nokiaPhone.Year = year
//	fmt.Printf("I am Nokia %d, I can call you!", nokiaPhone.Year)
//	return "Success!"
//}
//
//func (nokiaPhone NokiaPhone) callSelf() string {
//	fmt.Printf("I am Nokia %d, I can call you!", nokiaPhone.Year)
//	return "Success!"
//}
//
//
//type IPhone struct {
//}
//
//func (iPhone IPhone) call(year int) string {
//	fmt.Println("I am iPhone, I can call you!")
//	return "Success!"
//}
//
//func (iPhone IPhone) callSelf() string {
//	fmt.Println("I am iPhone, I can call you!")
//	return "Success!"
//}
//
//func main() {
//	var phone Phone
//
//	phone = new(NokiaPhone)
//	fmt.Println(phone.call(2010))
//	//fmt.Println(phone.Year)
//	phone = NokiaPhone{2012}
//	fmt.Println(phone.callSelf())
//
//	phone = new(IPhone)
//	phone.call(2021)
//
//}
