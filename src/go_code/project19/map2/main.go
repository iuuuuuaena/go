// map切片

package main


import "fmt"


func main(){

	// map切片
	// 首先声明
	var map1 []map[string]string
	// 然后给切片map分配内存空间
	map1 = make([]map[string]string,2)
	// 如果map切片的第一项为空，
	var count *int  = new(int)
	*count =1
	i :=fmt.Sprintf("%d",*count)
	if map1[0] == nil {
		// 就分配空间
		map1[0] = make(map[string]string , 2)
		map1[0][i] = "孙悟空"
		map1[0]["age"] = "550"
	}
	fmt.Println(map1)
	// [map[age:550 name:孙悟空] map[]]

	// 因为我们首先给map切片分配了两个内存空间，所以如果我们
	// 存入三个map的话，就会越界

	// 所以要动态增加

	//这样
	newMap := map[string]string{
		"name":"红孩儿",
		"age":"1000" ,
    }
	map1 = append(map1,newMap)
	fmt.Println(map1)
	//   [map[age:550 name:孙悟空] map[] map[age:1000 name:红孩儿]]
	// 这样我们就把map加入到了map切片中，动态增加了


}