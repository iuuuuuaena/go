package main

// 导入unsafe包，里面有Sizeof函数
import (
	"fmt";
	"unsafe"
)
 
/** 
	这个函数用来练习bool类型
*/
func main(){

	// bool类型
	// 只能是    true    和    false
	// 占用一个字节

	var a bool  =  true

	fmt.Printf("a = %b,类型是 %T,大小是 %d",a,a,unsafe.Sizeof(a))



}
