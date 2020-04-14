//  用于练习for语句

// 1.写一个输出矩形、棱形 、空心棱形的函数

package main

import (
	"fmt"
)
// 绘制矩形
func print1(num int){

	for i := 0;i<num;i++{
		for j := 0;j<num;j++{
			fmt.Printf("*")
		}
		fmt.Println("\n")
	}
	
}
// 绘制空心矩形
func print2(num int){

	for i := 0 ;i < num; i++{
		if i == 0||i == num-1{
			for j:=0;j<num;j++{
				fmt.Printf("*")
			}
		}else {
			fmt.Printf("*")
			for j:=1;j<num-1;j++{
				fmt.Printf(" ")
			}
			fmt.Printf("*")
		}
		fmt.Println("\n")
	}
}
// 绘制菱形
func print3(num int){

	for  i:= 0;i<num/2;i++{
		for j:=(num/2)-1-i;j>0;j--{
			fmt.Printf(" ")
		}
		for j:=0;j<2*i+1;j++{
			fmt.Printf("*")
		}
		fmt.Println("\n")
	}
	for  i:= 0;i<num/2;i++{
		for j:=0;j<i;j++{
			fmt.Printf(" ")
		}
		for j:=num-2*i;j>0;j--{
			fmt.Printf("*")
		}
		fmt.Println("\n")
	}
}
// 绘制空心菱形
func print4(num int){

	for  i:= 0;i<num/2;i++{
		for j:=(num/2)-1-i;j>0;j--{
			fmt.Printf(" ")
		}
	
		fmt.Printf("*")
		for j:=1;j<2*i;j++{	
			fmt.Printf(" ")
		}
		fmt.Printf("*")
		
		fmt.Println("\n")
	}
	for  i:= 0;i<num/2;i++{
		for j:=0;j<i;j++{
			fmt.Printf(" ")
		}

		fmt.Printf("*")
	
		for j:=num-2*i-1;j>1;j--{
			fmt.Printf(" ")
		}
		fmt.Printf("*")
		fmt.Println("\n")
	}
}

// 输出99乘法表
func printf(){

	for i := 1;i <= 9; i++{
		for j := 1 ; j <= i; j++{
			fmt.Printf("%dx%d=%d\t",j,i,i*j)
		}
		fmt.Println("\n")
	}

}

func main(){

	printf()

}