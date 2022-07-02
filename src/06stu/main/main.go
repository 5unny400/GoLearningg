package main

import "fmt"

func main() {
	var age int = 18
	//输出变量内存的地址  &:取地址符
	fmt.Println(&age) //0xc00000a088

	//定义一个指针变量  存储age变量地址的变量
	var ptr *int = &age
	fmt.Println(ptr)                //0xc00000a088
	fmt.Println("输出指针ptr的地址", &ptr) //0xc0000ce020

	//获取指针指向变量的值
	fmt.Println("获取指针指向变量的值:%v", *ptr)
}
