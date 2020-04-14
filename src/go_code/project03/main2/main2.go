package main


import "fmt"


func main(){


	// 7.变量可以被重新赋值，但是不能改变其类型

	var  i int  = 10

	fmt.Println("i = ",i)
	

	// 8.变量不可以被重定义在一个作用域里


	// 9、明确下，golang里面的变量
	// 必须声明，赋值，要不就不能通过编译的步骤

	// 10. +号的使用

	// (1) 相加
	var n1 = 10
	var n2 = 20
	var n3 = n1 + n2
	fmt.Println("n3 = n1 + n2 = ",n3)


	// (2) 字符串拼接

	var n4 string = "abc"
	var n5 string = "def"

	var n6 string = n4 + n5

	fmt.Println("n6是n4和n5的拼接，n6 = ",n6)

	/* 输出结果为：

		n3 = n1 + n2 =  30
		n6是n4和n5的拼接，n6 =  abcdef

	*/
}