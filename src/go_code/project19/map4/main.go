
// map的细节
package main

import "fmt"

func testMap(temp map[int]int){
	temp[1] = 100
}

//定义一个结构体
type myStruct struct {
	Name string
	Age int 
	Gender string
}

func main(){
	// 1.map是引用类型

	var map1 map[int]int
	map1 = make(map[int]int,10)

	map1[1] = 10
	map1[2] = 100
	map1[3] = 1000
	testMap(map1)

	fmt.Printf("map1[0]的地址为%p\n",&map1)
	// map的value是不可寻址的！！！！

	fmt.Printf("map1[1]的值为%d",map1[1])
	// map1[1]的值为100%   所以他是引用类型


	// map可以自动扩容


	// map的value经常使用结构体来存储


	map2 := make(map[string]myStruct,10)


	stu1 := myStruct{"tom",13,"gril"}
	stu2 := myStruct{"jarry",20,"boy"}

	map2["one"] = stu1
	map2["two"] = stu2


	fmt.Println("\n",map2)

	// map[one:{tom 13 gril} two:{jarry 20 boy}]

	//遍历

	for k,v := range map2{
		fmt.Printf("\n%s :\n \t %v\t%v\t%v",k,v.Name,v.Age,v.Gender)
	}

	type Users struct{
		U string
		P string
	}
	Users("123" ,"123")




}






