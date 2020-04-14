package main
//

import "fmt"

type Persion struct{
	Name string
	Age int
	Gender string
}
func main(){

	/*定义一个结构体，并给其赋值*/
 	persion1 :=  struct{
		Name string
		Age int
		Gender string
	}{"小明",100,"boy"}


	fmt.Println(persion1)
	

	fmt.Println("---")
}