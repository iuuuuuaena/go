package main

import "fmt"

// 排序
func main(){

	arr := [7]int{3,5,1,2,8,7,9}
	
	fmt.Println("arr = ",arr)
	// 传入一个指针
	buttleSort(&arr)
	fmt.Println("arr = ",arr)

	//arr =  [3 5 1 2 8 7 9]
	//arr =  [1 2 3 5 7 8 9]
	
	
}
/**
冒泡排序

*/
func buttleSort(arr *[7]int) {

	for i:= 0;i<len(*arr)-1;i++{
		for j:=0;j<len(*arr)-1-i;j++{
			if (*arr)[j] > (*arr)[j+1]{
				temp:=(*arr)[j]
				(*arr)[j]= (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}


}