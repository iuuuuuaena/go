package main

// 此文件用于练习输入语句


import (
	"fmt"

)

func main(){


	// 输入用户信息
	// 判断是否及格

	var name string
	var age byte
	var salary float32
	var isPass bool


	// scanln() : 一行一行输入

	fmt.Printf("请输入姓名:\n")
	//传一个地址进去，引用传递
	fmt.Scanln(&name)
	fmt.Printf("请输入年龄:\n")
	//传一个地址进去，引用传递
	fmt.Scanln(&age)
	fmt.Printf("请输入薪水:\n")
	//传一个地址进去，引用传递
	fmt.Scanln(&salary)

	if age < 40 && salary < 20000{
		isPass = true
		fmt.Printf("%v,通过测试",isPass)
	}else {
		isPass = false
		fmt.Printf("%v,没有通过测试",isPass)
	}

	// scanf():一次性输入

	fmt.Printf("请输入姓名，年龄，薪水，按空格隔开:\n")

	fmt.Scanf(" %s %d %f ",&name,&age,&salary)

	fmt.Printf("\n输入的是 name = %s,age = %d ,salary = %f",name,age,salary)


}