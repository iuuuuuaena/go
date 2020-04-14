
package main

import "fmt"
func main(){

	// 字典的学习

	// map

	// 1.声明


	var temp map[string]string
	// temp["三国演义"]= "刘备"
// ！！！
	fmt.Println(temp)
	// map[]  :为空！！！，不可直接赋值
	// 不能给空map赋值，会报错
	// panic: assignment to entry in nil map

	// goroutine 1 [running]:
	// main.main()
	//         /Users/mac/go/src/go_code/project19/map1/main.go:15 +0x4f
	// exit status 2

	// 所以给map赋值的时候，就是没有内存空间，必须先给他赋值

	// 用make 给map赋值

	//这句话的意思，是给map分配了10个空间！！！！
	temp = make(map[string]string,10)
	temp["三国演义"]= "刘备"

	fmt.Println(temp)

	// map[三国演义:刘备]
	// 同时，我们也看到了，第二种直接给map声明的方式

	b:=make(map[string]string,10)

	fmt.Println(b)
	//map[]
	//测试下大小
	fmt.Println(len(b))
	fmt.Println(len(temp))
	// 	0
	//  1   因为只有一组键值对！！！1


	// 赋值之前，一定要make

	// map的key不可以重复，但是value可以

	// map是无序的

	// 第三种声明方式

	var temp3 map[string]string = map[string]string{
		"西游记" : "孙悟空",
	}

	fmt.Println(temp3)

	temp4 := map [string]string{
		"西游记" : "孙悟空",
	}
	fmt.Println(temp3["西游记"] == temp4["西游记"])


	// 还可以 map[string ]map[string]string

	temp5 := make(map[string]map[string]string,10)

	temp5["class1"] = make(map[string]string)
	temp5["class1"]["name"]="tom"
	temp5["class1"]["age"]="11"
	temp5["class1"]["sex"]="boy"
	temp5["class2"] = make(map[string]string)
	temp5["class2"]["name"]="jarry"
	temp5["class2"]["age"]="22"
	temp5["class2"]["sex"]="gril"

	fmt.Println(temp5)
	// map[class1:map[age:11 name:tom sex:boy] class2:map[age:22 name:jarry sex:gril]]


	// map的增删改查！！！

	// 增加不用我说了！！

	// 删除用delete

	// 就把class1一个map删了
	delete(temp5,"class1")
	fmt.Println(temp5)
	// map[class2:map[age:22 name:jarry sex:gril]]

	// 一次性删除map，直接用make到一个新的map就行

	//再加上，一会用
	temp5["class1"] = make(map[string]string)
	temp5["class1"]["name"]="tom"
	temp5["class1"]["age"]="11"
	temp5["class1"]["sex"]="boy"

	// map的查找

	// OK是bool，如果有val，返回true，没有返回false
	val ,OK := temp5["class2"]

	if OK{
		fmt.Printf("\n有，值为%s\n",val)
	}else{
		fmt.Printf("\n没有\n")
	}

 
	// map的遍历  ，使用 for range

	for K,V:= range temp5{
		fmt.Printf("%s :\n",K)
		for k,v := range V{
			fmt.Printf("   %s = %s\n",k,v)
		}
		fmt.Println()
	}
	//class2 :
	//   age = 22
	//   sex = gril
	//   name = jarry

	// class1 :
	//   name = tom
	//   age = 11
	//   sex = boy

	
	



}

