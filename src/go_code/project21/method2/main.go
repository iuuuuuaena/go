// 方法！2
package main

import "fmt"

type A struct{

	Num int
}

func (a A) testMethod(){
	fmt.Println("test! num = ",a.Num)
}
func (a A) getSum(number int) int{
	sum := 0
	for i := 0;i<number;i++{
		sum +=number
	}

	return sum
}
// 这里传入的是指针，也就是引用传递
func (a *A) getSum2(num *int) int{
	sum := 0
	for i := 0;i<a.Num ;i++{
		sum +=(*num)
	}

	return sum
}
//自定义数据类型都可以有方法，不仅仅是struct
func main(){


	var a A
	num := a.getSum(50)
	// 求出1到50 的和
	fmt.Println(num)
	a.Num = 50
	num2 := a.getSum2(&a.Num)
	fmt.Println(num2)
// 	2500
//  2500

// 所以用结构题指针传递，的话，就是传入一个地址，可以直接对地址里面的值进行修改

// 所以无论是用指针传入方法，还是改变结构体的参数值，都可以输出结果，因为传递的都是一个地址！！！
}