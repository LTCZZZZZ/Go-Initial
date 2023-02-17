package variable

import (
	"log"
	"testing"
)

type TestS struct {
	a int
}

func TestChangeV1(t *testing.T) {

	T1 := &TestS{1}
	T2 := &*T1 // 这个时候先计算*T1得到值，然后&*T1取到地址，实际还是T1指向的地址
	T2.a = 5
	log.Print(T1)
	log.Print(T2)
}

func TestChangeV2(t *testing.T) {

	T1 := &TestS{1}
	T2 := *T1 // *T1的值赋给T2，T2是一个新的变量，不是T1指向的变量
	T3 := &T2
	T3.a = 5 // 所以这个改动不会影响T1
	log.Print(T1)
	log.Print(T2)
	log.Print(T3)
}
