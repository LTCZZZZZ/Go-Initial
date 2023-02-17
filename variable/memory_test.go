package variable

import (
	"fmt"
	"runtime"
	"testing"
	"unsafe"
)

type Person struct {
	Name    string
	Age     int
	Address *Address
}

type Address struct {
	City    string
	Country string
	Message string
}

func TestMemory(t *testing.T) {
	p := &Person{
		Name: "Alice",
		Age:  20,
		Address: &Address{
			City:    "Shanghai",
			Country: "China",
			Message: "HelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHello" +
				"HelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHelloHello",
		},
	}

	// 计算 p 本身占用的内存大小
	fmt.Println("Sizeof(p) =", unsafe.Sizeof(p))
	fmt.Println("Sizeof(*p) =", unsafe.Sizeof(*p))

	// 获取当前内存使用情况
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// 计算堆对象的大小
	fmt.Println("HeapAlloc =", m.HeapAlloc, "bytes")

	size := int64(unsafe.Sizeof(*p)) + int64(unsafe.Sizeof(*p.Address))
	fmt.Println("Sizeof(p)+Sizeof(p.Address) =", size, "bytes")

	// 计算 p 及其关联变量占用的内存大小（错误）
	fmt.Println("AllocSize(p) =", unsafe.Alignof(p))

	//// 获取当前内存使用情况
	//runtime.ReadMemStats(&m)
	//// 计算堆对象的大小
	//fmt.Println("HeapAlloc =", m.HeapAlloc, "bytes")
}
