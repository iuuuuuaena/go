package main

import "fmt"

/**
	使用这个函数来测试变量的声明
*/
func main(){
	// 1。如果只是声明变量，不赋值，int的默认值为0
	var i int

	fmt.Println("i =",i)
	//默认输出为0


	// 2.根据值自行判断类型
	var num = 11.02

	fmt.Println("num = ",num)


	// 3.还可以省略var，声明的同时赋值，可以私用：=
	// （temp := "" ）等同于 （var temp string   temp  = ""） 等同于这两句

	temp := "tom & jarry"

	fmt.Println("temp = ",temp)

	// 4.一次性声明多个变量	

	var n1, n2, n3 int

	fmt.Println("n1 = ",n1,"\tn2 = ",n2 ,"\tn3 = ",n3)

	// 5.同时声明不同类型的变量
	var n4, n5, n6  = 1 ,100,"你好"
	//当然也可以简写为:   n4, n5, n6  := 1 ,100,"你好"

	fmt.Println("n4 = ",n4,"\tn5 = ",n5 ,"\tn6= ",n6)


	// 6.还可以一次性声明多个变量


	var(
		n7 = 100
		n8 = "tom"
		n9 = 10.20
	)
	fmt.Println("n7 = ",n7,"\tn8 = ",n8 ,"\tn9= ",n9)
	
}