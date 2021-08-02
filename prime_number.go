package main

import (
	"fmt"
	"time"
)

func main() {
	// 从执行时间可以看出，1最慢，2好很多，3好非常多，4和3差不多
	t1 := time.Now()

	res1 := prime1(100000)
	fmt.Println(res1[:20])
	fmt.Println(time.Since(t1))

	t2 := time.Now()
	res2 := prime2(100000)
	fmt.Println(res2[:20])
	fmt.Println(time.Since(t2))

	t3 := time.Now()
	res3 := prime3(100000)
	fmt.Println(res3[:20])
	fmt.Println(time.Since(t3))

	t4 := time.Now()
	res4 := prime4(100000)
	fmt.Println(res4[:20])
	fmt.Println(time.Since(t4))

}

func prime1(a int) []int {
	var prime_list []int
	for i := 2; i < a; i++ {
		flag := 1
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = 0
				break
			}
		}
		if flag == 1 {
			prime_list = append(prime_list, i)
		}
	}
	return prime_list
}

func prime2(a int) []int {
	var prime_list []int
	for i := 2; i < a; i++ {
		flag := 1
		for _, j := range prime_list {
			if i%j == 0 {
				flag = 0
				break
			}
		}
		if flag == 1 {
			prime_list = append(prime_list, i)
		}
	}
	return prime_list
}

func prime3(a int) []int {
	var prime_list []int
	for i := 2; i < a; i++ {
		flag := 1
		for _, j := range prime_list {
			// 注意这里是大于，不能是大于等于
			if j*j > i {
				break
			}
			if i%j == 0 {
				flag = 0
				break
			}
		}
		if flag == 1 {
			prime_list = append(prime_list, i)
		}
	}
	return prime_list
}

func prime4(a int) []int {
	// 因为下面的循环所以这里初始条件略有不同
	prime_list := []int{2}
	for i := 3; i < a; i++ {
		flag := 1
		for j := 0; ; j++ {
			num := prime_list[j]
			// 注意这里是大于，不能是大于等于
			if num*num > i {
				break
			}
			if i%num == 0 {
				flag = 0
				break
			}
		}
		if flag == 1 {
			prime_list = append(prime_list, i)
		}
	}
	return prime_list
}
