package main

// 导入unsafe包，里面有Sizeof函数
import (
	"fmt";
	"unsafe"
)
 
/** 
	这个函数用来练习字符串类型 string
*/
func main(){


	// 基本使用	
	// 双引号
	// 默认占用16字节	
	// 使用utf-8编码
	// 字符串不可变的

	var a  string = "北京长城"

	fmt.Printf("a = %s ,类型是 %T,大小是%d",a,a,unsafe.Sizeof(a))


	// 之前说的字符串拼接用 +

	var b string = "abc"

	var c string = "def"

	var e = b + c 



	// 如果要跨行拼接，就需要 把+号写在后面


	f := "hello" + "hello" + "hello" + "hello" +
	"hello" + "hello" + "hello" + "hello"+
	"hello" + "hello" 
	// 这样 才不是错的

	



}