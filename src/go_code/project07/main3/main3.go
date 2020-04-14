// 此函数用于赋值的练习

package main

import (
	"fmt"
)

//测试函数
func test() int {
	
	return 100
}
func main() {



	a := 9
	b := 8
	fmt.Printf("交换前，a = %v ,b = %v",a,b)
	// 交换的操作
	t := a
	a = b
	b = t
	fmt.Printf("\n交换后，a = %v ,b = %v",a,b)


	a += 15 

	fmt.Printf("\na加15后为 = %v ",a)


	// 等号左边只能是变量 ，右边可以是变量，常量，函数，表达式

	d := test()


	fmt.Printf("\nd = test() = %v",d)

	/* 输出
			交换前，a = 9 ,b = 8
			交换后，a = 8 ,b = 9
			a加15后为 = 23 
			d = test() = 100%  

	*/

	// 测试： 交换a b ，但是不能用中间变量

	var x int = 10
	var y int = 100
	fmt.Printf("\n第二此交换前，a = %v ,b = %v",x,y)
	x = x + y
	y = x - y
	x = x - y
	fmt.Printf("\n第二此交换后，a = %v ,b = %v",x,y)

	// 输出结果：
	// 第二此交换前，a = 10 ,b = 100
	// 第二此交换后，a = 100 ,b = 10%  

}