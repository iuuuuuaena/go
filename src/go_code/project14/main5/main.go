package main

import(
	"fmt"
)

func main(){


	// 1.new一个* int类型的指针 （int ,float,struct)
	num := new(int)

	fmt.Printf("num的类型：%T,num的值 ：%v，num的地址为：%v",num,*num,num)
	// num的类型：*int,num的值 ：0，num的地址为：0xc000014080%   


	// 也可以用make分配引用类型的值

}