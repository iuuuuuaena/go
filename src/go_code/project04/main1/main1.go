//此文件用于探究类型转化
package main

import "fmt"


func main(){

	//类型转化
	//可以从低精度到高精度
	//也可以从高精度到低精度
	
	var a int  = 128

	var b  float32 = float32(a)

	var c float64 = float64(b)

	fmt.Printf("a = %d,type = %T , b = %f,type = %T, c = %f,type = %T",a,a,b,b,c,c)


	//如果从高精度转化成低精度，就会损失经度

	//然后按溢出处理，比如int8只能有-128-127，
	//如果你把128强制转化为int8，他的输出结果为-128，你细品，你自己的品，溢出处理

	var d int8 = int8(a)


	fmt.Printf("\nd = %d,type = %T",d,d)


	// go这里没有自动类型转化
	// 所以下面会报错

	var e int32 = 100
	var f int64 = 1
	var g int8 = 10
	// 这句会报错
	// e = f + 99
	// 错误为 
	// /main1.go:40:4: cannot use f + 99 (type int64) as type int32 in assignment

	//因为e是int64   f + 99是int32

	// 解决办法 ： 转换类型
	// e = int64(f) + 99

	// 这样编译能过
	g = int8(e) + 127 
	// 这样编译不能过,因为int8的0上只有127，你给他个127可以，但是128不行
	g = int8(e) + 128
	


	

}