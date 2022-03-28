package test

import (
	"fmt"
	"math"
	"testing"
	"time"
)

/// 测试等于和大于的运行速度
func TestEqualGt(t *testing.T) {

	m := 200000000
	n := int(math.Pow10(9))
	start_time := time.Now().UnixNano()
	fmt.Println(start_time)

	s1 := 0
	for i := 0; i < n; i++ {
		// 这里无论是用等于还是不等于，时间基本没差别，说明加法基本不耗时？？？
		// 此外，if语句中如果无内容，根据执行时间来看，这个判断本身都没有执行
		if i != m {
			s1 += i
		}
	}
	end_time1 := time.Now().UnixNano()

	s2 := 0
	for i := 0; i < n; i++ {
		if i > m {
			s2 += i
		}
	}
	end_time2 := time.Now().UnixNano()

	duration1 := end_time1 - start_time
	duration2 := end_time2 - end_time1
	fmt.Println(s1)
	fmt.Println(s2)
	//fmt.Println(duration1)
	fmt.Println(float64(duration1) / 1000000000)
	fmt.Println(float64(duration2) / 1000000000)
}
