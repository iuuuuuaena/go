package main


import "fmt"

// 接口


// 一个接口体可以实现多个接口

type mover interface{
	moveTo(where string)
}
type eater interface{
	eat(food string)
}

type cat struct{}


func (c *cat )eat(food string){
	fmt.Println("猫吃",food)
}


func (c *cat )move(place string){
	fmt.Printf("去%s旅游\n",place)
}


func main(){


	// 这个猫实现了两个接口的方法，所以他实现了两个接口



	//空接口

	// interface{}
	//空接口什么都能存放


	// 所以一般用于map


	var m1 map[string]interface{}


	m1 = make(map[string]interface{},16)

	m1["name"] = "tom"

	m1["age"] = 100
	
	m1["genfer"] = "boy"

	// 数组类型，...的意思是，让他自己跟具后面的个数算大小。
	m1["hobby"] = [...]string{"唱","跳","rap"}

	fmt.Println(m1)

	// map[age:100 genfer:boy hobby:[唱 跳 rap] name:tom]



	
	

}