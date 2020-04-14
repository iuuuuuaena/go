package main


import (
	"fmt"
	"math/rand"
	"time"
)

// 我们输入一个0-100之间的数，产生随机数，如果产生100次这个随机数，程序结束
func rannnn(num int) int{

	var count int
	// time.Now().Unix() :返回从1970年到现在的任意一时间
	// fmt.Println(time.Now().Unix())
	
	for {

		rand.Seed(time.Now().Unix())

		n := rand.Intn(100)+1
		// 为什么要加1 ，因为 生成的是[0,100）
		count++
		if n == num{
			break
		}
	}
	return count
}

func main(){

	var num int
	fmt.Println("请输入一个随机数：（0-100）")
	fmt.Scanln(&num)
	num = rannnn(num)
	
	fmt.Printf("一共用了%d次",num)

}


