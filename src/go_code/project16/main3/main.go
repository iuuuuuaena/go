package main

// 切片的操作函数

import (
	"fmt"
)

func main() {

	// 定义一个切片
	var r1 []int = []int{100, 200, 300}
	// 1.对切片进行追加同类型的元素

	// 追加，在返回去,底层是先创建一个新的数组，然后在把切片的首地址赋值给之前的切片首地址，首地址就更换了
	r1 = append(r1, 400)

	// 2.对切片追加另一个切片
	r2 := []int{121, 234, 2345}

	// 三个点是固定写法
	r1 = append(r1, r2...)

	fmt.Println(r1)
	// [100 200 300 400 121 234 2345]

	//3.切片的拷贝
	var rune1 []int = []int{1, 2, 3, 4, 5}

	var rune2 = make([]int, 10)
	//[0,0,0,0,0,0,0,0,0,0]

	// 使用copy函数，对切牌进行赋值，结果为
	copy(rune2, rune1)
	fmt.Println(rune2)

	//结果为[1 2 3 4 5 0 0 0 0 0]

	// 注意

	// 3.如果新切片的长度比旧的小，只拷贝对应的元素，并不会报错

	var rune3 = make([]int, 2)
	copy(rune3, rune1)
	fmt.Println(rune3)
	// [1 2]

	// 4.还有一种赋值操作

	var rune4 []int

	rune4 = rune2[:]

	fmt.Println("rune4=", rune4)
	// rune4= [1 2 3 4 5 0 0 0 0 0]

	// 5.因为切片是按引用传递，所以，会影响实参

	var rune5 = []int{1, 2, 3, 4, 5}
	fmt.Println("之前rune5=", rune5)
	// 前rune5= [1 2 3 4 5]
	testRune1(rune5)
	fmt.Println("之后rune5=", rune5)
	//后rune5= [100 2 3 4 5]

}

func testRune1(rune1 []int) {
	rune1[0] = 100
}
