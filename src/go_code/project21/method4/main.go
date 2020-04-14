// 方法和函数的区别
package main

import "fmt"


type Persion struct{

	Name string
}

func (p Persion)test01(){
	p.Name = "tom"
	fmt.Println("test01 , Name = ",p.Name)
}

func (p *Persion)test02(){
	// 因为在方法里，不用刻意强调指针，所以 （*p)和p一样，但是我建议还是写吧
	(*p).Name = "tom"
	fmt.Println("test02 , Name = ",(*p).Name)
}



func main(){

	// 1.调用方法不一样，



	// 2.传递参数时，限制不一样


	// 函数传入参数的时候，函数参数是什么，就得传什么，是值，就只能传值，传指针，就只能传地址进去

	// 但是方法，都可以

	// 如果是方法传入的是值传递，改变里面不会影响外面，
	// 但是方法传入的是结构体的指针，就会影响 了


	persion:= Persion{"!234"}

	persion.test01()
	fmt.Println("test01 , Name = ",persion.Name)
	// test01 , Name =  tom
	// test01 , Name =  !234
	// 因为是值拷贝，所以在test01里面改变了值，并不会影响外面的


	// 但是我用指针传递

	persion.test02()
	fmt.Println("test02 , Name = ",persion.Name)
	// test01 , Name =  tom
	// test01 , Name =  !234
	// test02 , Name =  tom
	// test02 , Name =  tom
	// 看，变了，对吧



	// 要注意，用&p调用test01函数，并不是引用传递，还是值传递

	(&persion).test01()
	fmt.Println("&&&test01 , Name = ",persion.Name)
	// 这里不会改变！！！！1





}