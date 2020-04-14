// 递归练习

package main


import "fmt"

// 斐波那契数

// 什么是斐波那契数呢，斐波那契数
// 就是， 从第三位开始每一个数都是前两个数之和
// 比如
// 1 1 2 3 5 8 13 21 34 65 99 
func test(num int ) int {
	if num == 1 || num == 2{
		return 1
	}else {
		return test(num-1)+test(num-2)
	}

}


// 猴子吃桃问题，他每天吃的都是前一天的一半再多一个,他十天后，只剩一个了，问一开始他有几个

func test2(day int) int {
	if day > 10 || day <1 {
		fmt.Println("对不起，数值不合标准，第十一天就吃完了")
		return 1
	}else {
		return (test2(day+1)+1)*2
	}
}



func main(){
	res := test(10)
	fmt.Println("5的斐波那契数为= ",res)

	var day int 
	fmt.Println("你想看第几天的桃子数")
	fmt.Scanln(&day)

	res2 := test2(day)

	fmt.Println(res2)
}