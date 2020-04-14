// 方法！！1
package main

import "fmt"

type A struct{

	Name string
}

func (a A) testMethod(){
	fmt.Println("test! name = ",a.Name)
}

//自定义数据类型都可以有方法，不仅仅是struct
func main(){

	// 方法！！！，结构体的方法

	// 方式是这样

	var a A 
	a.testMethod()
	// test! name =   因为我们还没有给他赋值，所以就是”“
	a.Name = "tom"
	a.testMethod()
	// test! name =  tom

	// testMethod，只能被结构题A的对象使用，其他结构体的对象使用会报错！！！1

	// 重点：方法传递是值传递，在里面改变name，不会对外面的值有所变化


	





}