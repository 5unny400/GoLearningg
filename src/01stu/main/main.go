package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(math.Pow(1, 1)) //指数运算
	fmt.Println(math.Max(1, 1)) //最大值

	var age int
	age = 19
	fmt.Println("age=", age)

	fmt.Println("----------------------------------------")
	//定义一个整数类型
	var num1 int8 = 120
	fmt.Println(num1)
	var num2 uint8 = 200
	fmt.Println(num2)

	var num3 = 28
	fmt.Printf("num3的数据类型是：%T\n", num3)
	fmt.Println("num3的字节数：", unsafe.Sizeof(num3))

	//表示学生年龄
	var agee uint8 = 18
	var ageee byte = 18
	fmt.Println(agee)
	fmt.Println(ageee)

}
