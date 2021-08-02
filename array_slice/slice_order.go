package main

import (
	"fmt"
	"sort"
)

func main() {
	// 倒序：
	var kArray = []string{"apt", "src", "fmt", "zoo", "amd", "yes"}
	fmt.Println("x" >= "y")

	// 稳定排序请使用SliceStable
	sort.Slice(kArray, func(i, j int) bool {
		return kArray[i] > kArray[j]
	})
	fmt.Println("逆序：", kArray)
	// 正序：
	sort.Strings(kArray)
	fmt.Println("正序：", kArray)
}
