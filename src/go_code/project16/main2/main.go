package main

// 切片的使用！！！
import (
	"fmt"
)

func main() {
	// 1、切片是slice

	// 2.切片是一个引用数据类型，是数组的地址

	// 3.使用与数组基本一致,长度动态变化的数组

	// 4、切片长度可以变化

	// 5. 直接写[]不用加长度

	// 1.创建一个数组
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	// 2.第二种方式，创建切片
	var rune2 = make([]int, 10)
	// 声明了一个长度为10 的切片，默人值为0[0 0 0 0 0 0 0 0 0 0]

	fmt.Println(rune2)

	// 引用了这个数组的 2到3这个位置，不包含三，（起始下表为1，终点是3，不到3）
	slice := arr[1:3]
	//  [2 3]

	fmt.Println(slice)

	fmt.Printf("切片的长度为：%d\n", len(slice))
	// 容量一般是长度的两倍
	fmt.Printf("切片的容量为：%d\n", cap(slice))

	//slice第一个元素存放的是arr的第二个元素的地址  slice[0] = &arr[1]
	//slice第二个元素存放的是切片的长度   2
	//slice第一个元素存放的是容量大小     4

	fmt.Printf("slice 第一个元素存放的地址为：%d", slice[0])
	fmt.Printf("\nslice 第一个元素的地址的值为：%p", &slice)
	fmt.Printf("\nslice 第一个元素存放的的地址的值为：%p", &slice[0])
	fmt.Printf("\nsarr的第二个值的地址的值为：%p", &arr[1])

}
