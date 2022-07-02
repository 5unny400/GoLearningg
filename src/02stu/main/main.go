package main

import "fmt"

func main() {
	var num1 float32 = 3.14
	fmt.Println(num1)
	var num2 float32 = -3.14
	fmt.Println(num2)
	var num3 float32 = 314e+2
	fmt.Println(num3)
	var num4 float32 = 314e+2
	fmt.Println(num4)
	var num5 float32 = 314e-2
	fmt.Println(num5)
	var num6 float32 = 314e-2
	fmt.Println(num6)

	var num7 float32 = 256.0000000012
	fmt.Println(num7)
	var num8 float64 = 256.0000000012
	fmt.Println(num8)

	//默认的浮点类型是float64
	var num9 = 3.17
	fmt.Printf("%T", num9)
}
