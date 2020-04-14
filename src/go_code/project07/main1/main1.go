// 此函数用于运算符的练习

package main1

import (
	"fmt"
)

func main(){
	// 六种运算符

	// 1.算数运算符  + -

	var a int = 1
	var b int = 2
	c := a + b
	fmt.Printf("c = a + b = %d\n",c)
	d := b - a
	fmt.Printf("d = b - a = %d\n",d)

	// /
	e := b / a  
	fmt.Printf("e = b / a = %d\n",e)

	var f int64 = 10
	var g int64 = 3

	// % 
	h := f % g
	fmt.Printf("h = f 除 g = %v \n",h)

	
	fmt.Printf("10 除 3 = %v \n",10 % 3)
	fmt.Printf("-10 除 3 = %v \n",-10 % 3)
	fmt.Printf("10 除 -3 = %v \n",10 % -3)
	fmt.Printf("-10 除 -3 = %v \n",-10 % -3)

	//  ++ 和 -- 

	var k = 10
	k++
	fmt.Printf("k++  = %v \n",k)
	k--
	fmt.Printf("k--  = %v \n",k)
	

	// 1.刚才需要注意的是 ，int类型在运算 / 的时候，会自动忽略小数部分，除非用float32和float64
	// 2. % 取余运算不能是float类型

	// 3.在说下 ++ 和 -- 
		// 在golang 里面， ++ 和 -- 语句只能独立存在，不能用到语句中
		// 这与很多语言不同
		// 比如， 我可以 
			// i := 1
			// i++
		// 但是不可以
			// var a = i++
		// 这样是不可以的
		// 而且golang 中 只有 后++ ，没有前++ 前--
	//只能a++,没有--a

	// 来个话设置温度转摄氏度的小例子

	var huashi float32 = 132.6

	var sheshi float32 = 5.0/9*(huashi -100)

	fmt.Printf("%.2f 转化为摄氏度是 %.2f",huashi,sheshi)

	
	
	




}