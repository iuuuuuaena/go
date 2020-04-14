// go语言的oop
package main

import "fmt"

type Test struct {
	Name   string
	Age    int
	Gender float64
	ptr    *int
	slice  []int
	map1   map[string]string
}

func main() {

	// 结构体声明的基本语法

	/*

			type  结构体名字  struct{
		         结构体字段  类型
			}

			这个字段很有意思，mysql里的列也叫字段

			如果没有赋值，就是默认值

			基本数据类型的话默认值就是自己的默认值


			引用数据类型像指针，切片，map这种默认值就是nil


			而且slice和map必须要make后才可使用






	*/

	var test Test
	test.slice = make([]int, 10)
	test.slice[0] = 1
	test.slice[1] = 2

	test.map1 = make(map[string]string, 10)
	test.map1["孙悟空"] = "123"
	fmt.Println(test.map1)

	//结构体的第二种使用方法

	test1 := Test{}

	test1.Name = "123"


	// 第三种方式 ，指针new

	var test2 *Test = new(Test)

	(*test2).Name = "jarry"


	// 第四种方式 ，引用调用

	var test3 *Test =  &Test{}

	(*test3).Name = "mike"

	








}
