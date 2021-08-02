//星号(*) :表示 cron 表达式能匹配该字段的所有值。如在第5个字段使用星号(month)，表示每个月
//斜线(/):表示增长间隔，如第2个字段(minutes) 值是 3-59/15，表示每小时的第3分钟开始执行一次，之后 每隔 15 分钟执行一次（即 3（3+0*15）、18（3+1*15）、33（3+2*15）、48（3+3*15） 这些时间点执行），这里也可以表示为：3/15
//逗号(,):用于枚举值，如第6个字段值是 MON,WED,FRI，表示 星期一、三、五 执行
//连字号(-):表示一个范围，如第3个字段的值为 9-17 表示 9am 到 5pm 直接每个小时（包括9和17）
//问号(?):只用于 日(Day of month) 和 星期(Day of week)，表示不指定值，可以用于代替 *

package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

//主函数
func main() {
	cron2 := cron.New() //创建一个cron实例

	//执行定时任务（每5秒执行一次）
	//err:= cron2.AddFunc("10-20/5 * * * * *", print5)
	//err:= cron2.AddFunc("*/5 * * * * *", print5)
	//err:= cron2.AddFunc("@every 1h30m10s", print5)
	err:= cron2.AddFunc("@every 5s", print5) // 不在整数秒，而从start开始每5s执行一次，Start()开始时不会执行，而是在5s后才执行第一次
	if err!=nil{
		fmt.Println(err)
	}

	//启动/关闭
	fmt.Println("start", time.Now())
	cron2.Start()
	fmt.Println("started", time.Now())
	defer cron2.Stop()
	select {
	//查询语句，保持程序运行，在这里等同于for{}
	}
}

//执行函数
func print5()  {
	fmt.Println(time.Now())
}
