
// 自定义数据类型

package main

import (
	"fmt"
	"go_code/project13/utils"
)

func init(){
	fmt.Println("\nmain的init()-----!!!!")
}

func test() (sum int ,sub int ){
	return
}
// 可变参数的演示
// 把这些参数加起来
// 可变参数必须放在后面
func test2(n1 int ,args... int) int {
	// 遍历args
	sum := n1
	// 把这些参数加起来
	for i := 0;i<len(args);i++{
		sum +=args[i]
	}
	return sum
}

func main(){


	//相当于给int 起了一个别名
	// 虽然把int
	// 当场了myint ，但是他们还是不能相互赋值 ，需要显性转换
	type myint int

	var a myint = 10
	fmt.Println("a = ",a)

	// 可以用 _ 忽略参数
	sum , _ :=test()

	fmt.Println("\nsum = ",sum)

	//还有可变参数
	// 把这些参数加起来
	abc := test2(20,20,1,1)
	fmt.Println("\nabc= ",abc)



	// init 函数 ()会在main之前，被go编译器调用

	// 通常可以用来做一些初始化的作用
	// 每一个源文件都可以有一个init函数

	// 变量定义是先执行  全局变量定义    》   init()函数   》   main函数


	// 每个包里都有init函数，比如我要引用utils包里的变量，我就可以先在

	// utils包的init函数中初始化，然后在我的main方法里调用

	// 那么究竟是utils包里的init函数先初始化呢还是，main函数里的init函数先初始化！！？？？？


	// 答案是 utils的init先初始化

	fmt.Println("Age = ",utils.Age,"Name = ",utils.Name)

	// 输出结果是

	/*
		uitls的init()-----!!!!

		main的init()-----!!!!
		a =  10

		sum =  0

		abc=  42
		Age =  15 Name =  Tom


	*/


	/// 匿名函数
	res := func (n1 int ,n2 int )int {
		return n1+n2
	}(10,20)
	fmt.Println("res = ",res)



	// 将func赋值给c ，c就是函数变量
	c := func (n1 int ,n2 int )int {
		return n1+n2
	}
	// 之后就可以通过c直接调用这个函数了
	res2 := c(10,20)
	fmt.Println("匿名函数res2 = ",res2)



	// 全局匿名函数 


	// 闭包！！！！！！！1
	// 累加器

	f := AddUpper()
	// 里面的n初始化之后，不会再初始化，
	fmt.Println(f(1))  // 11
	fmt.Println(f(2))  // 13
	fmt.Println(f(3))  // 16



}
// 累加器
// 这是个函数，名字叫ADDUpper ，返回一个函数
func AddUpper () func (int )int {


// 下面这五行，形成一个闭包，然后，里面的匿名函数，引用调用n，所以改变n，就会永远改变
	var n int =10
	return func(x int )int {
		n  += x
		return n
	}

	// 还有一个deffer 这几个东西 延时机制   释放资源 ,读文件什么的可以用dffer

}

func test3(){
	defer fmt.Println("OK1-----!!!")
	defer fmt.Println("OK2-----!!!")
	defer fmt.Println("OK3-----!!!")
	fmt.Println("OK4-----!!!")

	// 输出结果为  4   >   3   >2   >1   先进后出
}

// 可以用闭包做字符串拼接

// func AddUpper2()func(str string )string{
// 	return func() string{
// 		return "sd"
// 	}
// }