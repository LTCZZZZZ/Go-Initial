package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type OdsCourseRunningInfo struct {
	ZJID     string
	KSID     string
	KSSJKKSJ string
	KSSJJSSJ string
	SchId    uint32 // 学校ID，不上报
	InstId   uint32 // 教师来源学校或机构ID，不上报
	JSID     string
	JSLYLX   string
	XSID     string
	DKCNRDPJ string // 无评价则填"0"
	DKCJSDPJ string
}

//func (p OdsCourseRunningInfo) DataFormat() OdsCourseRunningInfo {
//	p.ZJID = strings.Join([]string{p.KSID, p.JSID, p.XSID}, "-")
//	return p
//}

func (p *OdsCourseRunningInfo) DataFormat(str string) {
	p.ZJID = strings.Join([]string{p.KSID, p.JSID, p.XSID, str}, "-")
}

//func TestReflect0(t *testing.T)  {
//	var p OdsCourseRunningInfo
//
//	v := reflect.ValueOf(p)
//	fmt.Println(v.NumField())
//	fmt.Println(v.NumMethod())
//	fmt.Println(v.ptr)
//	//v.ptr
//	fmt.Println(v.MethodByName("DataFormat"))
//	fmt.Println(reflect.ValueOf(&p).MethodByName("DataFormat"))
//}

func TestReflect(t *testing.T) {
	var sts interface{}
	sts = []OdsCourseRunningInfo{{}, {}}
	v := reflect.ValueOf(sts)
	for i := 0; i < v.Len(); i++ {
		//fmt.Println(v.Index(i))
		//s := v.Index(i)
		//fmt.Println(reflect.TypeOf(&s))
		//fmt.Println((&s).MethodByName("DataFormat"))
		fmt.Println(v.Index(i).Addr().MethodByName("DataFormat")) // 取到上面的指针方法

		// 最优解，使用指针方法
		//v.Index(i).Addr().MethodByName("DataFormat").Call(nil)
		v.Index(i).Addr().MethodByName("DataFormat").Call([]reflect.Value{reflect.ValueOf("test")})

		// 可行的方法一
		//v.Index(i).Set(v.Index(i).MethodByName("DataFormat").Call(nil)[0])

		// 可行的方法二，先转回原结构体
		//fmt.Println(v.Index(i).Interface().(OdsCourseRunningInfo).DataFormat())
	}
	fmt.Println(sts)
}
