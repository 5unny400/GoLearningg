package main

import "fmt"

func main() {
	var s1 string = "hello ,golang!"
	fmt.Println(s1)

	//字符串不可变，指的是字符串一旦被定义，其字符的值不可改变
	var s2 string = "abc"
	//s2[0] = 'c'  这样的操作是错误的
	fmt.Println(s2)

	//一般字符串
	var s3 string = "test"
	fmt.Println(s3)
	//特殊字符用反演符号 ` `包括起来
	var s4 string = `import "fmt"

func main(){
   //bool
   var flag01 bool = false
   fmt.Println(flag01)

   var flag02 bool = true
   fmt.Println(flag02)

   var flag03 bool = 5 < 9
   fmt.Println(flag03)
}`
	fmt.Println(s4)

	//字符串的拼接
	var s5 string = "abc" + "def"
	fmt.Println(s5)
	s5 += "hijk"
	fmt.Println(s5)

	//当一个字符串过长的时候,+保留在上一行的最后
	var s6 string = "abc" + "def" + "abc" + "def" + "abc" + "def" +
		"abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def" +
		"abc" + "def" + "abc" + "def" + "abc" + "def" + "abc" + "def"

	fmt.Println(s6)
}
