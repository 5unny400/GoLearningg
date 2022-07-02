package main //声明为main包，程序的入口包
import (
	"GoLearning/src/07stu/test"
	"fmt"
)

//程序入口函数
func main() {
	var age int = 18
	fmt.Println(age)
	fmt.Println(test.StuNo) //StuNo首字母在原包中必须大写，才能在包外引用

}
