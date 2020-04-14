package main

// 导入unsafe包，里面有Sizeof函数
import (
	"fmt",
	"unsafe"
)
 
/** 
	这个函数用来练习格式化输出
*/
func main(){


	// 1.使用 Printf格式化输出，%T 是可以输出变量的类型

	var a int  = 100
	var b string  = "absdasadasdc"
	
	fmt.Printf("a = %T\n",a)
	fmt.Printf("b = %T",b)


	// 2.使用unsafe包里的Sizeof函数，获得变量的大小

	fmt.Printf("\na的大小为 %d",unsafe.Sizeof(a))
	fmt.Printf("\nb的大小为 %d",unsafe.Sizeof(b))




}

