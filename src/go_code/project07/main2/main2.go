// 此函数用于运算符的练习

package main

import (
	"fmt"
)

//测试函数
func test() bool {
	fmt.Printf("\n test~!!!!!1")
	return true
}
func main() {

	//关系运算符
	// ==
	// >=
	// <=
	// !=
	// <
	// >
	// 不用掩饰了，太简单

	// 逻辑运算符

	// 逻辑与   && （短路与）
	// 逻辑或   || （短路或）
	// 逻辑非   ！

	a := 15
	if a < 10 && test() {
		fmt.Printf("\n测试短路与！！！！")
	}
	// 短路与，因为判断第一个条件为false，直接就把第二个条件忽略了
	fmt.Printf("\n程序结束")

	if a > 14 || test() {
		fmt.Printf("\n测试短路或！！！！")
	}
	// 短路或，因为判断第一个条件为true，直接就把第二个条件忽略了

}
