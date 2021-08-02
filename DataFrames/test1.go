package main

import (
	ctx "context"
	"fmt"
	dataframe "github.com/rocketlaunchr/dataframe-go"
)

func main() {
	s1 := dataframe.NewSeriesInt64("day", nil, 1, 2, 3, 4, 5, 6, 7, 8)
	s2 := dataframe.NewSeriesFloat64("sales", nil, 50.3, 23.4, 56.2, nil, nil, 84.2, 72, 89)
	df := dataframe.NewDataFrame(s1, s2)
	df.Append(nil, 9, 123.6)
	// 源码中只有添加map，没有添加结构体的选项
	df.Append(nil, map[string]interface{}{
		"day":   10,
		"sales": nil,
	})
	fmt.Println(df.NRows())
	fmt.Println(df.Names())
	fmt.Println(df.Row(0, false, dataframe.SeriesName))

	df.UpdateRow(0, nil, map[string]interface{}{
		"day":   3,
		"sales": 45,
	})
	df.UpdateRow(0, nil, 7, 77)

	sks := []dataframe.SortKey{
		{Key: "sales", Desc: true},
		{Key: "day", Desc: true},
	}

	// context包是Go 语言中用来设置截止日期、同步信号，传递请求相关值的结构体，是开发常用的并发控制技术。
	// 与WaitGroup的不同在于context可以控制多级的goroutine。
	// 排序
	df.Sort(ctx.Background(), sks)

	// 迭代
	iterator := df.ValuesIterator(dataframe.ValuesOptions{0, 1, true}) // Don't apply read lock because we are write locking from outside.

	df.Lock()
	for {
		row, vals, q := iterator(dataframe.SeriesName) // 如果不传参则returned value is a map containing the name of the series (string) and the index of the series (int) as keys.
		if row == nil {
			break
		}
		fmt.Println(*row, vals, q)
	}
	df.Unlock()

	fmt.Print(df.Table())

}
