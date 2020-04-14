package main

// 导入unsafe包，里面有Sizeof函数
import (
	"fmt";
	"unsafe"
)
 
/** 
	这个函数用来练习基本数据类型的默认值

*/
func main(){


	var a int  // 0

	var b float32 //0

	var c float64 //0

	var d bool //false

	var e string  // "" 空串

	fmt.Printf("a = %d 类型是 %T,b = %f ,c = %f ,d = %t , e = %v",a,unsafe.Sizeof(a),b,c,d,e)



	


}