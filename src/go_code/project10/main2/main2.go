// 测试 ：输入字母，然后改成大写 并输出


package main

import (
	"fmt"
)


func test(a byte){

	switch a{
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	case 'c':
		fmt.Println("C")
	case 'd':
		fmt.Println("D")
	case 'e':
		fmt.Println("E")
	default:
		fmt.Println("other")

	}

}

func test2(score int64){
	switch int64(score / 60){
	case 1:
		fmt.Println("A")
	case 2:
		fmt.Println("B")
	case 3:
		fmt.Println("C")
	default:
		fmt.Println("不及格")
	}
}

func main(){

	// var tem byte 
	// fmt.Println("请输入一个小写字母:")

	// fmt.Scanf("%c",&tem)

	// test(tem)
	var score int64
	fmt.Println("请输入一个分数")
	fmt.Scanln(&score)
	test2(score)
	fmt.Println("程序结束")
}