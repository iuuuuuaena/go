package main

// 导入unsafe包，里面有Sizeof函数
import (
	"fmt";
	"unsafe"
)
 
/** 
	这个函数用来练习字符类型
*/
func main(){
	// byte 和 string 的区别
	// byte使用'' 
	// string 使用"""
	var a byte = 1

	var b byte = 'a'

	// var c byte = "b"
	var c int = '北'
	
	fmt.Printf("\na = %c ，类型是 %T,大小是 %d",a,a,unsafe.Sizeof(a))
	fmt.Printf("\nb = %c ，类型是 %T,大小是 %d",b,b,unsafe.Sizeof(b))
	fmt.Printf("\nc = %c ，类型是 %T,大小是 %d，码值为 %d",c,c,unsafe.Sizeof(c),c)


	//总结
	/*
	一般般来讲，ascll中的字符，就用byte保存就行，但是不在ascll中的就不能用byte保存了
	就需要用大一点的类型保存，用int保存

	*/

	//go语言使用utf-8编码。
	// unicode包含 ascll码，是ascll的扩展
	//一个英文字母占用一个字节，一个汉字，占用三个字节

	var d int = 22269   //在 unicode中是  国这个汉字
	// 格式化输出为字符串是国
	fmt.Printf("\nd = %c",d)
	




}