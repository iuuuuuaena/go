package main


import (
	"fmt"
)

func main(){


	 //取地址符号 & 
	 // 取指针的值 *

	 var a int = 10

	 fmt.Printf("\na = %v, a 地址为 %p",a,&a)

	 var p *int = &a 

	 fmt.Printf("\np = &a ,p = %v,p 保存的值为%v, p 地址为 %p\n",&p,*p,&p)


	// 因为golang里面没有三元运算符

	// 所以只能用if代替

	var b = 30 
	var n int 
	if a > b {
		n = a
	}else {
		n = b
	}
	fmt.Printf("\nn = &v",n)

}