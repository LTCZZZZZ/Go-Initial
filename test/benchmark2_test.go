// go test -v -bench=Bench* benchmark2_test.go -benchmem
// B/op 值告诉您每个op分配了多少字节。 allocs/op 告诉每个操作发生了多少（不同的）内存分配
package test

import (
	"fmt"
	"testing"
	"unsafe"
)

func BenchmarkSlice100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getSlice(100)
	}
}
func BenchmarkSlice1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getSlice(1000)
	}
}
func BenchmarkSlice10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getSlice(10000)
	}
}
func BenchmarkMap100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getMap(100)
	}
}
func BenchmarkMap1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getMap(1000)
	}
}
func BenchmarkMap10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getMap(10000)
	}
}

func getSlice(size int) []int {
	//t := time.Now()
	s := make([]int, size*2)
	for i := 0; i < size; i++ {
		index := i << 1
		s[index] = i
		s[index+1] = i
	}
	//fmt.Println("slice time cost: ", time.Since(t))
	return s
}

func getMap(size int) map[int]int {
	//t := time.Now()
	m := make(map[int]int, size)
	for i := 0; i < size; i++ {
		m[i] = i
	}
	//fmt.Println("map time cost: ", time.Since(t))
	return m
}

// unsafe.SizeOf() 和 reflect.Type.Size() 仅返回传递值的大小，而不递归遍历数据结构并添加指向值的大小
func TestMem(t *testing.T) {
	size := 1000
	s := getSlice(size)
	m := getMap(size)
	fmt.Printf("slice size: %d\n", unsafe.Sizeof(s))
	fmt.Printf("map size: %d\n", unsafe.Sizeof(m))
}
