package main


import (
	"fmt"
	"time"
	"math/rand"
	
)
func main(){

	// 练习2：搞一个数组，然后插入一个数，插入到合适位置
	var insertValue int
	fmt.Scanf("%d",&insertValue)
	control(insertValue)

}
/*负责控制生成数组，控制排序，控制赋值*/
func control(n int){
	rand.Seed(time.Now().Unix())

	var arr [10]int = [10]int {}
	arrLen:= len(arr)

	for i:= 0;i<arrLen;i++{
		arr[i] = rand.Intn(100)
	}

	fmt.Println("排序之前的数组为-------\n")

	for i:=0;i<arrLen;i++{
		fmt.Printf("%d\t",arr[i])
	}
	quickSort(&arr,0,9)
	fmt.Println("\n排序之后的数组为-------\n")
	for i:=0;i<arrLen;i++{
		fmt.Printf("%d\t",arr[i])
	}

	var arr2 [11]int = [11]int {}
	var index int
	for i:=0;i<10;i++{
		if n < arr[i]{
			index = i
			break
		}else if n>arr[9]{
			index = 10
			break
		}
	}
	fmt.Printf("\n\nindex = %d\n",index)
	for i:=0;i<index;i++{
		arr2[i] = arr[i]
	}
	arr2[index] = n
	for i:=index;i<arrLen;i++{
		arr2[i+1] = arr[i]
	}
	fmt.Printf("\n把 %d 插入之后的数组为-------\n",n)
	for i:=0;i<arrLen+1;i++{
		fmt.Printf("%d\t",arr2[i])
	}

// 15      87      69      17      93      17      30      31      23      88      排序之后的数组为-------

// 15      17      17      23      30      31      69      87      88      93      index =  4

// 把 23 插入之后的数组为-------
// 15      17      17      23      23      30      31      69      87      88      93      %   


}
/*快速排序*/
func quickSort(arr *[10]int,left int ,right int){
	
	
	if left > right{
		return
	}
	l := left
	r := right
	// 找一个基准点
	lessValue:= (*arr)[left]

	for{
		if l < r {
			for {
				if l < r && lessValue <= (*arr)[r]{
					r--
				}else{
					break
				}
			}
			for {
				if l < r && lessValue >= (*arr)[l]{
					l++
				}else{
					break
				}
			}
			// 交换
			if l < r{
				temp:=(*arr)[r]
				(*arr)[r]=(*arr)[l]
				(*arr)[l] =temp			
			}
		}else {
			break
		}
		
	}
	//最后将基准位置与i和j相等的位置进行交换

	(*arr)[left] = (*arr)[l]
	(*arr)[l] = lessValue
	quickSort(arr,left,l-1)
	quickSort(arr,r+1,right)


}