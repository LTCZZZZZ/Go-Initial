package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 随机生成指定范围的浮点数
func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}

func main() {
	rand.Seed(time.Now().UnixNano())
	f := rand.Float64()
	fmt.Println(f + 100)
	fmt.Println(fmt.Sprintf("%.10f", f+100))
	fmt.Println(strconv.FormatFloat(f+100, 'f', 5, 64))
	fmt.Println(randFloats(1.10, 101.98, 5))
}
