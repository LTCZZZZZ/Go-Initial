// 当前程序的包名
package main

// 导入其他包
// 省略调用，调用的时候只需要Println()，而不需要fmt.Println()
import . "fmt"

// 常量定义
const PI = 3.14

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

// 全局变量的声明和赋值
var name = "gopher"

// 一般类型声明
type newType int

// 结构的声明
type gopher struct{}

// 接口的声明
type golang interface{}

// 由main函数作为程序入口点启动
func main() {
	Println(KB, MB)
	Println("Hello World!")
	s1 := `第一行
    第二行
    第三行
    ` // 反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。
	Println(s1)
}
