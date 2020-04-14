// 此文件主要用于练习错误处理


package main


import(
	"fmt"
)
/**
模拟在test函数出现了错误，回到main函数之后的程序如何执行
*/
func test(){

	// 使用defer + recover来处理异常
	defer func(){  // 匿名函数
		err := recover() // recover内置函数，可以捕获异常,还可以把这句话写到if后面，if后面可以加初始化的语句
		if err != nil{  // 说明捕获到异常
			fmt.Println("err = ",err)
			fmt.Println("将错误发送给邮箱管理员jxd先生~~",err)

			/* 输出结果
			err =  runtime error: integer divide by zero
			将错误发送给邮箱管理员jxd先生~~ runtime error: integer divide by zero
			main中的方法！！！！1

			*/  
		}
	}()
	
	// 有了这个错误处理，就可以捕获异常，程序就不会轻易挂掉

	num1 := 10
	num2 := 0
	fmt.Println("test res = ",num1/num2)
	fmt.Println()


}
func main(){


	test()

	fmt.Println("main中的方法！！！！1")

	// panic: runtime error: integer divide by zero

	// goroutine 1 [running]:
	// main.test()
	//         /Users/mac/go/src/go_code/project15/error/main.go:17 +0x11
	// main.main()
	//         /Users/mac/go/src/go_code/project15/error/main.go:26 +0x22
	// exit status 2
	/*

		错误处理机制
		发生panic时,程序会崩溃，可以捕获到错误，并进行处理，保证代码可以继续执行
		捕获到错误后，给管理员预警机制，可能是邮件或者是短信。


		golang 不用try-catch-finally这种形式

		golang 使用 defer  panic  和 recover

	*/

	// 使用dffer + recover来处理异常
}