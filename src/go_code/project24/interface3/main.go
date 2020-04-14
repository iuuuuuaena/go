package main

import "fmt"

// 接口类型断言

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, OK := a.(string)
	if !OK{
		fmt.Println("猜错了")
	}else {
		fmt.Println("猜对了，是",str)
	}
	fmt.Println(str)
}

// 接口类型断言2

func assign2(a interface{}){

	switch a.(type){
	case string:
		fmt.Println("是一个字符串：",a.(string))
	case int:
		fmt.Println("是一个int：",a.(int))
	case bool:
		fmt.Println("是一个bool：",a.(bool))
	case float32:
		fmt.Println("是一个浮点数：",a.(float32))
	}
}
func main() {


	var b float32 = 12.0
	assign2(b)


}
