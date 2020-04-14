//此文件主要用于学习go语言的转义字符
//跟c中的一样

package main

import "fmt"

//fmt中主要提供格式化和输入输出

func main() {
	// \n换行符
	fmt.Println("abcd\n1234 \n\n")
	// \\输出一个斜杠
	fmt.Println("1\\2\\3\\4\\5\\ \n")
	// \r换行，后面的会把光标移到这行的前面，并把之前的相应位置覆盖
	//所以下面只会输出 你是个屁第一，把之前的四个覆盖了
	fmt.Println("我是宇宙第一\r你是个屁\n")
	// \t 制表符
	fmt.Println("1\t2\t3\t4\t5\t\n")

	/* 输出结果
	abcd
	1234


	1\2\3\4\5\

	你是个屁第一

	1       2       3       4       5




	*/

}
