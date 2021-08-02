// 重要 important
package main

import (
	"fmt"
)

func main() {
	var cities = [5]string{}
	var sliceMap = make(map[string]*[]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		//temp := make([]string, 0, 3)
		temp := cities[0:3:5]
		value = &temp
	}
	//value = append(value, "北京", "上海")
	(*value)[0] = "北京"
	(*value)[1] = "上海"
	sliceMap[key] = value
	fmt.Println(sliceMap, *(sliceMap[key]))

	*value = []string{"后海"} // 生效
	fmt.Println(sliceMap, *(sliceMap[key]))

	//fmt.Println(sliceMap.B)  // 怎样访问结构体hmap中的属性？？？

	p := &(*sliceMap[key])[0]
	//fmt.Printf("p地址：%p\n", p)
	*p = "后上海" // 生效
	fmt.Println(sliceMap, *(sliceMap[key]))

	// ——————————————————————————————————————————————————————————————————————————
	v := *sliceMap[key]
	fmt.Printf("%p\n", sliceMap[key])
	fmt.Printf("%p\n", &v) // 两地址不同，所以这个地方操作的时候需要换一种方式

	p = &v[0]
	*p = "后后海" // 生效
	//fmt.Printf("p地址：%p\n", p)
	v = append(v, "大上海") // 无效
	fmt.Printf("%p\n", &v)
	//v[1] = "大上海"
	fmt.Println(sliceMap, *(sliceMap[key]))

	// ——————————————————————————————————————————————————————————————————————————
	vp := sliceMap[key]
	fmt.Printf("%p\n", sliceMap[key])
	fmt.Printf("%p\n", vp) // 两地址相同

	p = &(*vp)[0]
	*p = "指针后海" // 生效
	//fmt.Printf("p地址：%p\n", p)
	*vp = append(*vp, "指针上海") // 生效相当于用*vp代替了上面的v
	fmt.Printf("%p\n", &v)
	//v[1] = "大上海"
	fmt.Println(sliceMap, *(sliceMap[key]))

	// 指数级自复制
	for i := 0; i < 2; i++ {
		*vp = append(*vp, *vp...)
	}
	fmt.Println(sliceMap, *(sliceMap[key]))
}

// 这里就有个问题，怎么取map的某key对应的value的地址？即 &map[key]，
// map的value本身是不可寻址的，因为map中的值会在内存中移动，并且旧的指针地址在map改变时会变得无效

// 如果实在需要类似的功能，将map的value直接设为指针类型
