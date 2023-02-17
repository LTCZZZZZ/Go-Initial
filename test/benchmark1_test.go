// go test -v -bench=Bench* benchmark1_test.go -benchmem
package test

import (
	"testing"
)

// 斐波那契数列
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
func sum(a, b int) int {
	return a + b
}

func BenchmarkFib10(b *testing.B) {
	//fmt.Println(b.N)
	for n := 0; n < b.N; n++ {
		fib(10)
	}
}
func BenchmarkFib20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(30)
	}
}
func BenchmarkSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sum(1, 2)
	}
}
