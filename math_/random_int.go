package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 根据我的理解，一旦随机数种子确定，其后每次获取随机数的结果也随之确定，只会根据次数不同而不同
	// rand相当于是一个对象，它内部是记录了你获取随机数的次数的
	// 此猜想已被time.Sleep(100 * time.Millisecond)这一行验证，
	// 即在不设置随机数种子时，无论sleep间隔多久，得到的结果集始终完全相同
	//fmt.Println(time.Now().UnixNano())
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))  // 如果不设置随机数种子，则每次执行结果相同
	for index := 0; index < 10; index++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(rand.Intn(10))
	}
}
