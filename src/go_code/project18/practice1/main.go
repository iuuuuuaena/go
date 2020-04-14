//练习1

package main


import(
	"fmt"
	"time"
	"math/rand"

)

func main(){

	//随机生成十个数（0-100）保存到，数组中，求他的平均值，并查找是否有55
	 
	var arr [10]int  = [10]int{}
	rand.Seed(time.Now().Unix())

	
	for i:=0;i<10;i++{
		arr[i] = rand.Intn(100)
	}
	for i:= 0;i<10;i++{
		fmt.Printf("%d\t",arr[i])
	}

	sum:=0
	for i:=0;i<len(arr);i++{
		sum += arr[i]
	}

	result:=sum/len(arr)

	fmt.Println("\n平均值为= ",result)


}