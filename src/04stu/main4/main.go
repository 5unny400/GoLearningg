package main

import "fmt"

func main() {
	var n int = 32
	fmt.Println(n)
	//var n2 float32 = n 这样是错误的
	//必须显式转换
	var n2 float32 = float32(n)
	fmt.Println(n2)
	//注意：n依旧是int类型，只是将n的值转为了float32
	fmt.Printf("%T\n", n) //输出n的类型并换行

	var n3 int64 = 88888
	var n4 int8 = int8(n3)
	fmt.Println("n3=", n3)
	fmt.Println("n4=", n4)

	var n5 int32 = 12
	//要匹配等号左右的数据类型
	var n6 int64 = int64(n5) + 30
	fmt.Println(n6)

	var n7 int64 = 12
	var n8 int8 = int8(n7) + 127
	//var n9 int8 = int8(n7)+128  编译报错

	fmt.Println("n8=", n8)
}
