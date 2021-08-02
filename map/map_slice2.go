package main

import (
	"fmt"
)

var cities = [5]string{}

func main() {
	//fmt.Println(cities)
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		//value = make([]string, 0, 3)
		value = cities[0:3:5]
	}
	//value = append(value, "北京", "上海")
	value[0] = "北京"
	value[1] = "上海"
	sliceMap[key] = value

	//value = sliceMap[key]
	//value = []string{"后海", }
	//fmt.Println(sliceMap)

	//fmt.Println(sliceMap.B)  // 怎样访问结构体hmap中的属性？？？

	//ptrs := new([]string)
	//fmt.Printf("ptrs地址：%p\n", ptrs)
	//*ptrs = sliceMap[key]
	//fmt.Printf("ptrs地址：%p\n", ptrs)  // 两地址相同，说明此操作不会有意义

	p := &sliceMap[key][0]
	//fmt.Printf("p地址：%p\n", p)
	*p = "后海"
	fmt.Println(sliceMap)

	v := sliceMap[key]
	//fmt.Printf("%p\n", &v)
	p = &v[0] // 这里修改的是slice对应的底层数组，所以有效
	//fmt.Printf("p地址：%p\n", p)
	*p = "后后海"
	v[1] = "大上海"
	fmt.Println(sliceMap)
	//fmt.Println(len(v), cap(v))

	//fmt.Printf("append前：%p\n", &v)
	v = append(v, "天津")
	//v[2] = "天津"
	//fmt.Printf("append后：%p\n", &v) // 前后地址没变
	fmt.Println(sliceMap)

	v2 := sliceMap[key] // 这里获取到的slice已经不是map中原始的slice了，但它和原slice对应相同的底层数组array
	vp := &v2
	//fmt.Printf("%p", vp)  // 可以看到这里获取到的地址和上面的&v不同，
	fmt.Println(*vp)
	*vp = []string{"前海", "后海"}
	//fmt.Printf("%p", vp)
	fmt.Println(*vp)
	//fmt.Printf("p地址：%p\n", &sliceMap[key][0])
	fmt.Println(sliceMap)
}

// 这里就有个问题，怎么取map的某key对应的value的地址？即 &map[key]，
// map的value本身是不可寻址的，因为map中的值会在内存中移动，并且旧的指针地址在map改变时会变得无效

// 如果实在需要类似的功能，将map的value直接设为指针类型，参见map_slice3_2
