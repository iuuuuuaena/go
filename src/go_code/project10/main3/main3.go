// 练习循环

package main

import("fmt")


func test(num int16) (int16 ,int8){

	var sum int16
	var count int8
	var i int16 = 0
	for  ;i < num ; i++ {
		sum += i + 1
		count++
	}
	return sum,count

}
func main(){

	// 输入一个数，输出从1到他的和 和数量
	var score int16
	
	var sum int16 = 0

	var count int8 = 0
	fmt.Println("请输入一个数：")
	fmt.Scanf("%d",&score)
	sum ,count = test(score)

	fmt.Printf("sum = %d,count = %d",sum,count)


	fmt.Println("程序结束")


}