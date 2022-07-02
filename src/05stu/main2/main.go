package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n1 int = 19
	var s1 string = strconv.FormatInt(int64(n1), 10) //第一个参数必须为int64,第二个参数指定什么进制
	fmt.Printf("s1对应的类型：%T,s1=%q\n", s1, s1)

	var n2 float32 = 2.34
	//第一个参数为float64 第二个参数为输出格式 第三个参数为小数点后的数字各位为-1时采用尽可能少的小数点后数  第四个参数表示第一个参数的来源类型
	var s2 string = strconv.FormatFloat(float64(n2), 'f', 9, 32)
	fmt.Printf("s2对应的类型：%T,s2=%q\n", s2, s2)

	var n3 bool = true
	var s3 string = strconv.FormatBool(n3)
	fmt.Printf("s3对应的类型：%T,s3=%q\n", s3, s3)

}
