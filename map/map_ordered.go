package main

import (
	"fmt"
	"sort"
)

func main() {
	map1 := make(map[int]string, 5)
	map1[1] = "www.topgoer.com"
	map1[2] = "rpc.topgoer.com"
	map1[5] = "ceshi"
	map1[3] = "xiaohong"
	map1[4] = "xiaohuang"
	sli := []int{}
	for k := range map1 {
		sli = append(sli, k)
	}
	//sort.Ints(sli)  // 正序
	sort.Slice(sli, func(i, j int) bool {
		return sli[i] > sli[j] // 倒序
	})
	for i := 0; i < len(map1); i++ {
		fmt.Println(sli[i], map1[sli[i]])
	}
}
