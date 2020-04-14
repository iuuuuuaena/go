package main

import "fmt"

func main(){


	// 使用二分查找的方式！！！1
	arr := [8]int{1,3,4,8,9,10,35,63}
	retu := banrSearch1(&arr,10)
	fmt.Println("rturn=",retu)
	banrSearch2(&arr,0,7,10)
	
}


/*第一种使用for循环*/
func banrSearch1(arr *[8]int ,target int)(index int){

	// if len(*arr) == 0{
	// 	return -1
	// }
	left := 0
	right :=len(*arr)-1
	fmt.Println("len = ",len(*arr))
	for{
		if left > right{
			break
		}
		middle := (right+left)/2
		if target > (*arr)[middle]{
			left = middle+1
		}else if target < (*arr)[middle]{
			right = middle -1
		}else {
			return middle
		}
	}
	return -1

}


/*第二种使用递归*/
func banrSearch2(arr *[8]int,left int,right int ,target int){

	if left >right{
		fmt.Println("找不到")
		return
	}
	middle := (left+right)/2
	if target > (*arr)[middle] {
        banrSearch2(arr,middle+1,right,target)
	}else if target < (*arr)[middle]{
		banrSearch2(arr,left,middle-1,target)
	}else{
		fmt.Println("retu2=",middle)
		
	}


}