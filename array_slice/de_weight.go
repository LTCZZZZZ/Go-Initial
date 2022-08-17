// 一维数组去重
package main

import "fmt"

func main() {
	list := []string {"hello", "hello"}
	list = removeRepeatElement(list)
	fmt.Println(list)
}

func removeRepeatElement(list []string) []string{
	// 创建一个临时map用来存储数组元素
	temp := make(map[string]bool)
	index := 0
	for _, v := range list {
		// 遍历数组元素，判断此元素是否已经存在map中
		_, ok := temp[v]
		if ok {
			list = append(list[:index], list[index+1:]...)
		} else {
			temp[v] = true
		}
		index++
	}
	return list
}
// 输出：
// [hello]
