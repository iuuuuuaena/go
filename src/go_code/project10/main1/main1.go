// 流程控制语句


package main

import(
	"fmt"
	// "math"
)

func main(){

// 	// 练习流程控制


// 	//1. 顺序

	

// 	//2. 分支

// 	var name string
// 	var age byte
// 	var salary float32
// 	var isPass bool


// 	// scanln() : 一行一行输入

// 	fmt.Printf("请输入姓名:\n")
// 	//传一个地址进去，引用传递
// 	fmt.Scanln(&name)
// 	fmt.Printf("请输入年龄:\n")
// 	//传一个地址进去，引用传递
// 	fmt.Scanln(&age)
// 	fmt.Printf("请输入薪水:\n")
// 	//传一个地址进去，引用传递
// 	fmt.Scanln(&salary)
// 	// 不用写括号
// 	if age < 40 && salary < 20000{
// 		isPass = true
// 		fmt.Printf("%v,通过测试",isPass)
// 	}else {
// 		isPass = false
// 		fmt.Printf("%v,没有通过测试",isPass)
// 	}

// 	// golang支持在if语句中直接定义一个变量

// 	 if i := 30 ;i > 20{
// 		 fmt.Println("我是无敌的")
// 	 }
// 	 // 并且，golang里面，不允许不带{} ，尽管你可能只有一句话，但是也得带
 

	 // 多分支
	 var a int = 3
	 if a == 1{
		fmt.Println("a = ")
	 }else if a == 2{
		fmt.Println("我不认识")
	 }else if a == 3{
		fmt.Println("我不认识")
	 }else {
		 fmt.Println("我不认识")
	 }
	 // 注意 ：在golang里面，不能再条件上赋值
//	下面这个是不允许的 ，只能在外面赋值，然后拿到这里比较
	//  if a = 5{

	//  }

// 	b := math.Sqrt(float64(a))
// 	fmt.Println("b = ",b)


// 	// switch 

// 	// case 唯一，不用带break

// 	fmt.Println("请输入一个数a：\n")
// 	var tem int 
// 	fmt.Scanln(&tem)
// 	// switch后面可以带函数返回的值
// 	// case可以带多个值,并且类型要跟switch那个匹配
// 	switch  tem {
// 		case  1:
// 			fmt.Println("a== 1")
// 			fallthrough
// 		case  2:
// 			fmt.Println("a == 2")
// 		case  3,4,5:
// 			fmt.Println("a == 3")
// 		default :
// 			fmt.Println("不知道")

// 	}
// 	// fallthrough 穿透，直接执行下个语句，不考虑tem == 2



	//3. 循环

	// for 循环
	//打印多句话

	for i:= 0 ;i < 10 ; i++{
		fmt.Println("你好！")
	}


	// 字符串遍历

	 var str string =  "sadsas asas"
 
	for i:= 0 ; i <len(str) ; i++{
		fmt.Printf("%c",str[i])
	}


	fmt.Println("\n")
	// for range

	for index ,val := range str{
		fmt.Printf("%d,%c\t",index,val)
	}

	// 因为汉字在golang里占三个字节
	// 用字符，输出，就只输出了一个字节，是乱码

	// 解决方式是用切片


	var str2 string =  "sadsas我是好人"

	str3 := []rune(str2)
	for i:= 0 ; i <len(str3) ; i++{
		fmt.Printf("%c \n",str3[i])
	}

	// 这样就不会是乱麻了

	// 还有就是，用for range 的话，是不会有乱码的问题的，因为for range是以字符串输出的
	for index ,val := range str2{
		fmt.Printf("%d,%c\t",index,val)
	}


// for 循环还可以省略里面的步骤 
// 下面的这个相当于 ;; 
	var i = 0
	for {
		if i < 10{
			break
		}
		fmt.Println(i+1)
		i++
	}

}
