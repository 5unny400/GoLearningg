package main

import "fmt"

func main() {
	var c1 byte = 'a'
	fmt.Println(c1)

	var c2 byte = '6'
	fmt.Println(c2)

	var c3 byte = '('
	fmt.Println(c3)

	//	查看UTF-8编码表:
	//http://www.mytju.com/classcode/tools/encode_utf8.asp
	//汉字底层对应的Unicode码
	var c4 int = '中' //对应的整型为20013
	fmt.Println(c4)
	//golang 的字符对应的使用的是utf-8编码（utf-8是Unicode的一种编码方案）

	var c5 byte = 'A'
	fmt.Printf("c5对应的字符为：%C", c5)
}
