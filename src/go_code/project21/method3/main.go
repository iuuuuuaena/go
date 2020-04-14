// 方法！2
package main

import "fmt"

type Person struct{
	Num int
}

func(person *Person) getPoint(){

	fmt.Printf("2person.num的地址为%p",&person.Num)
}

func main(){

	var person Person
	person.Num = 5
	fmt.Printf("1peeso.num的地址为%p\n",&person.Num)
	person.getPoint()
// 可以看到，是一样的
	// 1peeso.num的地址为0xc000014080
	// 2person.num的地址为0xc000014080% 
	
	// 传地址进方法中！！！
}