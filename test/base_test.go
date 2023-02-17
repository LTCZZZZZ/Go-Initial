package test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

type TeaInstId struct {
	TeacherId     uint32
	InstitutionId uint32
	Name          string
}

func TestBase(t *testing.T) {
	var index, index1, index2 []TeaInstId
	var name = "test"
	for i := 0; i < 10000; i++ {
		name += "1"
	}
	t1 := TeaInstId{TeacherId: 1, InstitutionId: 2, Name: name}
	index2 = append(index, t1)
	index = append(index1, index2...)
	//fmt.Println(index)
	fmt.Println("size name: ", unsafe.Sizeof(name))
	fmt.Println("size t1: ", unsafe.Sizeof(t1))
	fmt.Println("size index: ", unsafe.Sizeof(index))

	var (
		a uint32 = 1
		b uint32 = 2
	)
	fmt.Println(a - b)

	fmt.Println("a" + "bcd")

	fmt.Sprintf("%02d", 3)

	fmt.Println(time.Now().Unix())

	var teas []TeaInstId
	// 对slice类型make，实际上是在slice里面添加了len个数据
	teas = make([]TeaInstId, 3) // 会生成3个元素，每个元素都是零值
	fmt.Println(teas)

	teas = make([]TeaInstId, 0, 5) // 如果要设置容量，应该这么设置
	fmt.Println(teas)
	fmt.Println(len(teas), cap(teas))

	var err error
	err = errors.New("test error")
	fmt.Println(err.Error()) // error=nil时，调用Error()会panic

	c := 'A'
	fmt.Println(reflect.TypeOf(c))
	fmt.Printf("%c\n", c)
	fmt.Printf("%d\n", c)
}
