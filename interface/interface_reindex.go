package main

import (
	"MyTools"
	"fmt"
	"reflect"
)

func reindex(s []interface{}, col string, index []interface{}, item interface{}) []interface{} {
	// s是结构体切片，col是列名，res是返回的结构体切片，item是默认的对象的指针
	if len(s) == len(index) {
		return s
	}
	res := make([]interface{}, len(index))
	vItem := reflect.ValueOf(item).Elem()
	vItemCol := vItem.FieldByName(col)

	for num, i := range index {
		signal := false // 是否匹配到的标志符
		for _, j := range s {
			//if reflect.ValueOf(j).FieldByName(col) == reflect.ValueOf(i) {  // 这种方法不行，竟然不相等
			if reflect.ValueOf(j).FieldByName(col).Interface() == i {
				signal = true
				res[num] = j
				break
			}
		}
		if !signal { // 如果没匹配到
			fmt.Println(reflect.ValueOf(i))
			vItemCol.Set(reflect.ValueOf(i))
			res[num] = vItem.Interface()
		}
	}
	return res
}

type User struct {
	AddTime int
	Name    string
	Age     int
}

func main() {
	var s, index []interface{}
	s = append(s, User{3, "三", 30})
	fmt.Println(s)
	col := "AddTime"
	for _, i := range MyTools.Pyrange(1, 8) {
		index = append(index, i)
	}
	item := User{}
	res := reindex(s, col, index, &item)
	fmt.Println(res)

}
