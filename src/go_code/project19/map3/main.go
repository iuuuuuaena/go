// golang中的map没有排序方法


package main


import (
	"fmt"
	"sort"
)

func main(){
	//map排序

	map1 :=make(map[int]int,10)
	map1[1] = 1
	map1[2] = 2
	map1[3] = 6
	map1[4] = 9
	fmt.Println(map1)

	// 我我们可以看到。，map默认是无需的

	//排序

	// 1.放入slice
	// 2.对slice排序
	// 3.遍历slice，按照key值输出map

	var sliceKeys[]int 
	for k,_:= range map1{
		sliceKeys = append(sliceKeys,k)
	}
	fmt.Println(sliceKeys)

	fmt.Println("---------")

	//使用sort里的ints函数从小到大排序slice
	sort.Ints(sliceKeys)
	fmt.Println(sliceKeys)
	fmt.Println("---------")
	for _,v := range sliceKeys{
		fmt.Printf("%d = %d\t",v,map1[v])
	}
}