package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

// 结构体表示一个函数，他实现了Run接口，结构体的字段即为函数的参数
type Func struct {
	arg1 int
	arg2 int
	arg3 string
}

func (f Func) Run ()  {
	fmt.Println(time.Now())
	fmt.Println(f.arg1 + f.arg2)
	fmt.Println(f.arg3)
}

//主函数
func main() {
	c := cron.New() //创建一个cron实例
	c.AddJob("*/5 * * * * *", Func{2,3, "自定义参数"})
	c.Start()
	defer c.Stop()
	select {}
}