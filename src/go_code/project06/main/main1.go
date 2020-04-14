// 标识符命名规范
package main


import (
	"fmt"

	//引入model包
	"go_code/project06/model"
)

func main(){
		//26 大小写字母+_+0-9

		//区分大小写

		//标识符不能含有空格

		// _ 是空标识符，用于站位

		// 还有很多保留关键字，自己查

		//重点

		//  我们的包名一般来讲最好跟我们的这个文件所在的文件夹保持一致，这样很舒服
		// 比如我这个main1.go在project06里面，按理来讲，我就应该package project06 ，这样是最符合规范的
		// 当然我们还可以在project06下面创建一个main文件夹，然后把这个文件放进去，就可以了


		// 变量函数命名

		// 驼峰法

		// myFun()
		// toString()
		// changeType()

		// go语言里面，函数和变量首字母是大写的，就可以被其他文件访问，如果是小写，就不可以
		// 相当于， 大写开头的就是public， 小写开头的就是private


		// 引入其他包

		// 使用包里的变量

		// 导入首字母是大写的
		fmt.Println("TestVar = \n",model.TestVar)
		// 导入首字母是xiao写的
		// fmt.Println("testVar = ",model.testVar)

		// 会报错
		// command-line-arguments
		// ./main1.go:49:28: cannot refer to unexported name model.testVar
		// ./main1.go:49:28: undefined: model.testVar

		
	
}