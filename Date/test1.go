package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println(time.Unix(int64(1620576000), 0).Format("01-02"))
	fmt.Println(time.Unix(int64(1620576000), 0).In(time.FixedZone("UTC+8", 8*3600)).Format("01-02"))
}
