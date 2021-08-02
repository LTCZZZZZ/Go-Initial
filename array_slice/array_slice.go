package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 声明一个底层数组，长度为10，容量为10
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(reflect.TypeOf(arr))
	fmt.Printf("[%T]len(arr)=%d,cap(arr)=%d \n", arr, len(arr), cap(arr))
	// 声明两个切片，分别取底层数组的[1,4],[7:]
	s1 := arr[2:5]
	fmt.Printf("[%T]len(s1)=%d,cap(s1)=%d \n", s1, len(s1), cap(s1))
	s2 := arr[7:]
	fmt.Printf("[%T]len(s2)=%d,cap(s2)=%d \n", s2, len(s2), cap(s2))
	// 数组越界指的是超出数组长度,例如下面报：panic: runtime error: index out of range [3] with length 3
	//fmt.Println(s1[3])
	// s1可以添加元素，添加后s1长度变大，容量不变，同时底层数组被修改，这里TM就离谱，离谱！！！
	s3 := append(s1, 20)
	fmt.Println(s1)
	fmt.Println(s3)
	fmt.Printf("[%T]len(s1)=%d,cap(s1)=%d \n", s1, len(s1), cap(s1))
	fmt.Println(arr)
	// 上面可能的解释
	// 切片实际的是获取数组的某一部分，len切片<=cap切片<=len数组，
	// 切片由三部分组成：指向底层数组的指针、len、cap
	// s2添加元素时，由于容量还在其内不用扩展，所以是其指向的arr产生实际变化

	// s2如果需要添加元素，因为容量不够，需要进行扩容，开辟新数组，将原来的7,8,9拷贝过来，再添加一个20，长度为4，容量为6
	s2 = append(s2, 20)
	fmt.Printf("[%T]len(s2)=%d,cap(s2)=%d \n", s2, len(s2), cap(s2))
	// 此时修改s2的数组，底层数组arr不再受影响
	s2[1] = 10
	fmt.Println(arr)

	// 容量为什么为6？涉及到数组的扩容规则，举个例子如下：
	ints := []int{1, 2}          // 原容量oldCap =2
	ints = append(ints, 3, 4, 5) // 预估容量cap = 5
	/*
		 if oldCap * 2 < cap {
			newCap = cap
		} else {
			if oldLen < 1024 {
				newCap = oldCap *2
			}else if oldLen >= 1024 {
				newCap = oldCap *1.25
			}
		}
	*/
	// 上面例子中newCap = 5,int数组所占字节为5*8 = 40，但go语言向内存管理模块向操作系统申请的内存容量却没有40大小的，只有48符合，于是newCap = 48/8 = 6
	// go语言内存管理模块是16bytes叠加的，8，16，32，48，64，80，96
	// 参考博客：https://www.cnblogs.com/ldaniel/p/8502867.html?utm_source=debugrun&utm_medium=referral
	fmt.Printf("[%T]len(ints)=%d,cap(ints)=%d \n", ints, len(ints), cap(ints))
	// 例子验证第二种情况
	ints2 := []int{1, 2}
	ints2 = append(ints2, 3)
	// 此时oldCap * 2 > cap ,满足第二种情况，newCap = 4
	fmt.Printf("[%T]len(ints2)=%d,cap(ints2)=%d \n", ints2, len(ints2), cap(ints2))
}
