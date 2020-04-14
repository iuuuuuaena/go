// 工厂模式式
package main

import (
	"fmt"
	"go_code/project23/model"
)


func main(){


	stu1 := model.Student{
		Name:"tom",
		Score : 100,
	}
	fmt.Println(stu1)
	// {tom 100}
	
	// 如果我们的model包里的student结构体首字母是小写的怎么办，不可访问了就


	// person1 := model.persion{
	// 	Name :"jarryu",
	// 	Age :123
	// }

	// 上面这段代码，报错，因为persion是小写

	// 可以使用工厂模式，解决这个问题


	//  其实工厂模式就是用个函数，跟java的get函数一样，把对象返回而已，看model包


	 persion2 := model.NewPersion("jjj",1000)

	//  fmt.Println(persion2)
	//  fmt.Println((*persion2))
	//  fmt.Println((*persion2).Name）
	//  fmt.Println((*persion2).Age）
	//  // &{jjj 1000}
	//  // {jjj 1000}
	//  // jjj
	//  // 1000

	//完美解决

	//如果结构体里的字段都是小写的，怎么办
	// 就绑定几个方法返回即可

	// 来，看

	// 一样可以输出
	fmt.Println(persion2)
	fmt.Println((*persion2))
	fmt.Println((*persion2).GetName())
	fmt.Println((*persion2).GetAge())
	// &{jjj 1000}
	// {jjj 1000}
	// jjj
	// 1000


	





}