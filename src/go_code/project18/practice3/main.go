package main

import "fmt"

func main(){

	//编写一个 程序，让二维数组的四周都变为0

	var binArr[5][6]int = [5][6]int{}
	for i:=0;i<len(binArr);i++{
		for j:=0;j<len(binArr[0]);j++{
			binArr[i][j] = 1
		}
	}


	for i,k:= range binArr{
		for j,v := range k{
			fmt.Printf("b[%d][%d] = %d\t",i,j,v)
		}
		fmt.Println()
	}
	for i:=0;i<len(binArr[0]);i++{
		binArr[0][i] = 0
		binArr[len(binArr)-1][i] = 0
	}
	for i:=0;i<len(binArr);i++{
		binArr[i][0] = 0
		binArr[i][len(binArr[0])-1] = 0
	}
	for i,k:= range binArr{
		for j,v := range k{
			fmt.Printf("b[%d][%d] = %d\t",i,j,v)
		}
		fmt.Println()
	}
	
}