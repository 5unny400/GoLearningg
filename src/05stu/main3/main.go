package main

import (
	"fmt"
	"strconv"
)

//golang中文文档
//https://studygolang.com/pkgdoc
func main() {
	//string-->>bool
	var s1 string = "true"
	var b bool
	//ParseBool函数有两个返回值:(value bool,err error)
	//err用_忽略
	b, _ = strconv.ParseBool(s1)
	fmt.Printf("b对应的类型：%T,b=%v\n", b, b)

	//string-->>int64
	var s2 string = "19"
	var c int64
	c, _ = strconv.ParseInt(s2, 10, 64)
	fmt.Printf("c对应的类型：%T,c=%v\n", c, c)

	//string-->>float64
	var s3 string = "21.2"
	var d float64
	d, _ = strconv.ParseFloat(s3, 64)
	fmt.Printf("d对应的类型：%T,d=%v\n", d, d)

	//注意string转换为对应的基本类型应该具备相应的意义，否则转换后是对应基本类型的默认值
	var s4 string = "true"
	var e bool
	e, _ = strconv.ParseBool(s4)
	fmt.Printf("e对应的类型：%T,e=%v\n", e, e)
	var s5 string = "true"
	var f int64
	f, _ = strconv.ParseInt(s5, 10, 64)
	fmt.Printf("f对应的类型：%T,f=%v\n", f, f)
}
