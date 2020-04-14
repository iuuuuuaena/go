package main

import "fmt"

func main(){
	// 二维数组

    // 默认为0
	var arr [2][3]int 
	fmt.Println(arr)
	// [[0 0 0 0 0 0] [0 0 0 0 0 0] [0 0 0 0 0 0] [0 0 0 0 0 0]]


	// 遍历1
	// 二维数组在内存中的存储方式，如果这个int型的二维数组有两行三列，就申请两个指针，分别指向一个一维数组，每个一维数组是三个元素，
	// 如果是三行二列，这个二维数组就有三个指针，指向三个两个元素的一维数组，
	// 第一种情况，这两个指针之间的地址差距是24，当然是用16进制计算的，
	for i:= 0 ;i<len(arr);i++{
		for j:= 0;j<len(arr[0]);j++{
			fmt.Printf("%d\t",arr[i][j])
		}
		fmt.Println()
	}
	fmt.Println("--------------------------")
	// 遍历2
	// for range
	for i,k :=  range arr {
		for j,v :=  range k{
		 fmt.Printf("arr[%d][%d] = %d\t",i,j,v)
		}
		fmt.Println()
	}
	//arr[0][0] = 0   arr[0][1] = 0   arr[0][2] = 0
	//arr[1][0] = 0   arr[1][1] = 0   arr[1][2] = 0


	fmt.Printf("\n%p",&arr[0])
	fmt.Printf("\n%p\n",&arr[1])
	//
	// 0       0       0
	// 0       0       0

	// 0xc0000ac060
	// 0xc0000ac078%  

// 看这里，这里差24，三个int类型的块，这是个三个元素的一维数组


	// 二维数组的第二种初始化的方式

	var arr2[3][2] int = [3][2]int{{1,2},{3,1},{4,5}}
	fmt.Println(arr2)
	// [[1 2] [3 1] [4 5]]


	// 其他声明的方式就那几种，省略什么的

}