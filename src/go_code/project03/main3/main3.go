package main


import "fmt"


func main(){
	//测试整数类型
	/* 整数类型包括
		int   64位系统占8个字节，32位占4个 跟int64
		int8    负2的七次方---正的 2的七次方减一    1字节
		int16  负2的十五次方---正的2的十五次方减一   2字节
		int32									4字节
		int64                      				8字节
		uint8   0-255                           
		uint16  0到2的16次方减一
		uint32  0到2的32次方减一
		uint64  0到2的64次方减一

		rune   类似于int32 有unicode编码，适合储存中文     4字节
		byte   等价于 uint8                             1字节
		(1)不加u的是分正负，加u的是正的
		(2)为什么负的比正的多1，因为二进制用第一位表示符号位，0为正，1为负，
			如果是1000和0000这个两个数都代表0，这样就会出现两个0，这样是不行的
			所以我们把正的剪掉一，就可以避免
		(3)int 8范围是  -128~127



	*/


	var i int8 = -128

	fmt.Println("i = ",i)
	//可以正常输出
	//但是吧i改为-129，就会报错

	i = -129
	fmt.Println("i = ",i)
	//报错   36:4: constant -129 overflows int8


	// 无符号的整形也一样
	var j uint8 = 255

	fmt.Println("j = ",j)
	//可以正常输出
	//但是吧i改为256，就会报错

	j = 256
	fmt.Println("j = ",j)
	//报错   48:4: constant 256 overflows uint8

	var a uint  = 100

	// byte 相当于无符号int8
	var b byte = 255

	var c rune = 100000

	fmt.Println("a = ",a,"\tb = ",b,"\tc = ",c)


	//golang 默认整形为int
	

}
