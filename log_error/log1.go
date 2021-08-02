// log.Fatal 和 log.Panic 不只是 log
// Go 的 log 包在调用 Fatal*()、Panic*() 时能做更多日志外的事
package main

import "log"

func main() {
	log.Fatal("Fatal level log: log entry") // 输出信息后，程序终止执行
	log.Println("Nomal level log: log entry")
}
