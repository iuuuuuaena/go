// go语言的oop
package main

import "fmt"

type Cat struct {
	Name   string
	Age    int
	Gender string
}

func main() {

	// 1.go里面没有class
	// 2.go里面用结构体实现oop
	// 3.go里面没有继承的概念，但是他是通过其他方式实现继承的
	// 4.go里面没有方法重载，构造析构，this
	// 5.go还是有封装继承多态的概念的，只是跟其他语言不一样

	// 6.go天然支持oop，低耦合

	// 7.go面向接口编程！！！！！

	// 1.定义一个结构体

	// 2.声明一个对象
	var cat1 Cat
	fmt.Println(cat1)
	// 输出：{ 0 }  ,为什么呢？因为没有赋值，cat1用的是默认值 "" 0 "" 所以是{ 0 }

	// 3.赋值
	cat1.Name = "tom"
	cat1.Age = 12
	cat1.Gender = "boy"
	fmt.Println(cat1)
	// {tom 12 boy}

	fmt.Println("go的oop！！！")

	// go中的结构体是值传递，不是先指定一个地址，
	// 然后指向另一块内存的，而是直接指向一块内存

}
