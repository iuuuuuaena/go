package main

import "fmt"

func main(){


	// 编写一个菲薄那器，放入切片

	fmt.Println("fbn=",fbn(10))

	//fbn= [1 1 2 3 5 8 13 21 34 55]

}

func fbn(n int )( []uint64){

	// 1、声明一个切片

	var slice = make([]uint64,n)
	slice[0] =1 
	slice[1] =1 
	for i:=2 ;i<n;i++{
		slice[i]= slice[i-1]+slice[i-2]
	}
	return slice
	
}