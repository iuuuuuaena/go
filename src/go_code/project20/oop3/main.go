// go语言的oop
package main

import "fmt"

type Test struct {
	Name   string
	Age    int
	Gender float64
	ptr    *int
	slice  []int
	map1   map[string]string
}


type Point struct{
	x int 
	y int
}

type Triangle struct{
	left, right Point
}
func main() {

	var test1 Test


	test1.Name = "123"
	fmt.Println(test1.Name)

	// test2是指向test1的结构体指针1！！！！！！！
	var test2 *Test = &test1

	test2.Name = "456"

	fmt.Println(test2.Name)



	// 结构体的字段在内存中是连续分布的，

	// 1.

	var triangle Triangle



	fmt.Printf("\npoint地址为：%p，保存的地址为%p",&triangle.left,triangle.left)
	fmt.Printf("\npoint,x地址为：%p，保存的地址为%p",&triangle.left.x,triangle.left.x)

	// point地址为：0xc000018100，保存的地址为%!p(main.Point={0 0})
    // point,x地址为：0xc000018100，保存的地址为%!p(int=0)%  



	 //结构体里的两个指针地址是连续的，但是他们指向的地址不一定是连续的



	// 两个结构体之间的赋值

	// 可以使用强制类型转化，但是必须保证两个结构体的字段一致


	// 因为两个结构体是自己重新定义的类型，所以他们是不同的类型！！！！，不能直接相互赋值
}
