package main

// 导入unsafe包，里面有Sizeof函数
import (
	"fmt";
	"unsafe"
)
 
/** 
	这个函数用来练习浮点数
*/
func main(){


	//浮点数包括 
	/*
		float32   4字节
		float64   8字节

		浮点数都是有正负的
	*/

	var  a  float32 = 10.20
	var  b  float64 = -12939213.0

	fmt.Printf("\na = %f ，类型是 %T,大小是 %d",a,a,unsafe.Sizeof(a))
	fmt.Printf("\nb = %.10f ，类型是 %T,大小是 %d",b,b,unsafe.Sizeof(b))

	// golang默认浮点数是用的float64类型

	c  := 1.1

	fmt.Printf("\nc = %f ，默认类型是 %T,大小是 %d",c,c,unsafe.Sizeof(c))


	// 浮点型还可以省略小数点前的0 

	d := .1234

	fmt.Printf("\nd = %f ，类型是 %T,大小是 %d",d,d,unsafe.Sizeof(d))


	// 科学计数法

	e := 5.1234e3 //相当于 5123.4 

	fmt.Printf("\ne = %f ，类型是 %T,大小是 %d",e,e,unsafe.Sizeof(e))



}