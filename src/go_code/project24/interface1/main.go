package main

import "fmt"
// 接口

// 接口是一个类型，他不关什么什么类型传入，只关心方法
type speaker interface{
	// 把speak方法放入interface里面，就可以实现传入任意
	// 实现了speak方法的类型都是speaker类型
	speak()
}
type cat struct{}

type dog struct{}

func (c *cat)speak(){
	fmt.Println("miaomiao!")
}
func (d *dog)speak(){
	fmt.Println("wangwang!")
}


func do(x speaker){
	// 不管传入谁，我都想调用他的speak方法！！！
	x.speak()

}
//接口，不管传入什么类类型的参数，我都可以调用他的放方法
func main(){


	cat1 :=cat{}
	dog1 := dog{}
	do(&cat1)
	fmt.Println("----")
	do(&dog1)
	// miaomiao!
	// ----
	// wangwang!





	// 一个变量如果实现了一个接口里面的所有方法，这个变量就实现了这个接口，都可以称为这个接口的变量！



	// 接口是个类型，ss实现了speaker中的所有方法
	var ss speaker

	ss = &cat1
	fmt.Println(ss)
	// {}

	// 简单来讲，就是如果你一个对象实现了我接口里的所有方法，
	// 那我新建一个结构体类的变量，你们俩就可以相互赋值


	ss = &cat1
	// 看，我cat1是一个cat类型的对象，实现了speak方法，所以
	// cat1和ss是等价的


	// 使用值接受着，和指针接受着不一样


	

	

}