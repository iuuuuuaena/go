// 方法；练习

package main


import "fmt"

type Animal struct {
	Name string 
	Num int
}
// 给定一个数，判断是奇数还是偶数
func (animal *Animal) JudgeNum(){
	if animal.Num %2 == 0{
		fmt.Println("偶数")
	}else {
		fmt.Println("偶数")
	}
}

func main(){


	var aminal Animal
	aminal.Num = 11

	aminal.JudgeNum()

}