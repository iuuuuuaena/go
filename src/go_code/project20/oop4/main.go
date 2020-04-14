// go语言的oop
package main

import (
	"fmt"
	"encoding/json"
)
type Persion struct{

	Name string `json:"name"`
	Age int   `json:"age"`
	Gender string  `json:"gender"`
}


func main(){


	// go 的结构体中有个tag，是用来序列化和反序列化的
	// 将一个结构体序列化成json
	// 跟客户端进行信息传输


	 persion := Persion{"Tom",10,"gril"}
	 

	 fmt.Println("")


	 jsonStr , err := json.Marshal(persion)

	 if err != nil{
		 fmt.Println("json转换错误")
	 }else{
		 fmt.Println(jsonStr)
		// 这是序列化之后的
		//  [123 34 78 97 109 101 34 58 34 84 111 109 34 44 34 65 103 101 34 58 49 48 44 34 71 101 110 100 101 114 34 58 34 103 114 105 108 34 125]
	 }

	 fmt.Println(string(jsonStr))
	//  {"Name":"Tom","Age":10,"Gender":"gril"}
	// 这种首字母大写的字段名，其他程序员可能看不懂
	// 所以我们要把他改成小写，就用	`json:"name"` 这就叫加个tag


	//改过之后的结果为：{"name":"Tom","age":10,"gender":"gril"}




}