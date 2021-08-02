// golang 和 python 差不多，函数返回值直接写在参数后边。不过内部定义不能写函数名，调用可以通过变量来使用。
package main

func main() {
	i := 0
	println(i)
	change := func() {
		//i := 10  // 特别注意，加上这一行的话，i就是change函数内的局部变量，如果不加，i就是外部变量
		// 补充说明：说实话，没有nonlocal关键字，这个显得很不清晰
		i += 1
		println(i)
	}
	change()
	println(i)
}
