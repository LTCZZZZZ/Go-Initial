package main

import "fmt"

func main() {
	changeSliceTest()
}

func changeSliceTest() {
	arr1 := []int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	fmt.Printf("arr1: %T\narr2: %T\n", arr1, arr2)
	arr3 := [3]int{1, 2, 3}
	arr4 := []int{}
	arr4 = append(arr4, []int{3, 5, 7}...) // ...传参类似于python中的*，不过位置不同，一个在后一个在前
	fmt.Println(arr4)
	arr5 := make([]int, 10, 20)
	arr5[2] = 2
	fmt.Println(arr5)
	b := arr5[0:5]
	fmt.Println(len(arr5), cap(arr5), len(b), cap(b))
	a := make([]int, 1)
	b = append(a, 1)
	fmt.Println(len(a), cap(a), len(b), cap(b))
	_ = append(a, 2)
	fmt.Println(len(a), cap(a), len(b), cap(b))
	println(b[0], &a, &b)

	fmt.Println("before change arr1, ", arr1)
	changeSlice(arr1) // slice 按引用传递
	fmt.Println("after change arr1, ", arr1)

	fmt.Println("before change arr2, ", arr2)
	changeArray(arr2) // array 按值传递
	fmt.Println("after change arr2, ", arr2)

	fmt.Println("before change arr3, ", arr3)
	changeArrayByPointer(&arr3) // 可以显式取array的 指针
	fmt.Println("after change arr3, ", arr3)
}

func changeSlice(arr []int) {
	arr[0] = 9999
}

func changeArray(arr [3]int) {
	arr[0] = 6666
}

func changeArrayByPointer(arr *[3]int) {
	arr[0] = 6666
}
