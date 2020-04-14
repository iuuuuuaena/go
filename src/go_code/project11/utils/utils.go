
package utils

// 这个包是工具包，通常用来放一些自己写的功能函数
import "fmt"
// 函数 首字母要大写，说明是个静态的函数
// 输出99乘法表
func Out(){

	for i := 1;i <= 9; i++{
		for j := 1 ; j <= i; j++{
			fmt.Printf("%dx%d=%d\t",j,i,i*j)
		}
		fmt.Println("\n")
	}

}