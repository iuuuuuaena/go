// 基本数据类型和string进行转换

package main

import (
	"fmt";
	"unsafe";
	"strconv"
)

func main(){

	// 基本数据类型转string
	// 1.使用fmt中的Sprintf（）

	var a int = 99
	var b float64 = 100.1
	var c bool = true
	var d byte = 'n'
	var e string = ""

	e = fmt.Sprintf("%d",a)
	fmt.Printf("e = %v, type = %T,大小是%d\n",e,e,unsafe.Sizeof(e))

	e = fmt.Sprintf("%f",b)
	fmt.Printf("e = %v, type = %T\n",e,e)

	e = fmt.Sprintf("%v",c)
	fmt.Printf("e = %v, type = %T\n",e,e)

	e = fmt.Sprintf("%c",d)
	fmt.Printf("e = %v, type = %T\n",e,e)


	// 2.第二种方式，使用 strconv的formatint函数
	//FormatTP类函数将其它类型转string：FormatBool()、FormatFloat()、FormatInt()、FormatUint()
	// func FormatInt(i int64, base int) string
	// func FormatUint(i uint64, base int) string
	// 第二个参数base指定将第一个参数转换为多少进制，有效值为2<=base<=36。当指定的进制位大于10的时候，超出10的数值以a-z字母表示。例如16进制时，10-15的数字分别使用a-f表示，17进制时，10-16的数值分别使用a-g表示。

	// 例如：FormatInt(-42, 16)表示将-42转换为16进制数，转换的结果为-2a。
	e = strconv.FormatInt(int64(a),10)
	// %q是加引号的string
	fmt.Printf("使用strconv转换.e = %v\n",e)

	e = strconv.FormatFloat(b,'f',10,64)
	// b 要转的  f：格式  10：保留多少位  64,表示是float64
	fmt.Printf("使用strconv转换.e = %q\n",e)



	// 还可以用strconv中的 itoa，也是讲int转为string

	e = strconv.Itoa(a)

	




}
