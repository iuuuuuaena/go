package main
// 数组和切片的使用
import (
	"fmt"
	"math/rand"
	"time"
)
func main(){
	// 定义一个大小为5的
	var  arrs [5]int

	var p *int = &arrs[0]

	for i := 0 ;i<len(arrs);i++{
		fmt.Printf("arrs[%d] = %d,地址为：%p\n",i,arrs[i],&arrs[i])
	}
	fmt.Println(arrs)

	fmt.Printf("p保存的地址为：%p,p 的地址为=%p\n",p,&p)
	// [0 0 0 0 0]

	/* 没浪个地址之间，差的是int类型的字节数
 
	arrs[0] = 0,地址为：0xc0000ac060
	arrs[1] = 0,地址为：0xc0000ac068
	arrs[2] = 0,地址为：0xc0000ac070
	arrs[3] = 0,地址为：0xc0000ac078
	arrs[4] = 0,地址为：0xc0000ac080
	[0 0 0 0 0]
	p保存的地址为：0xc0000ac060,p 的地址为=0xc0000ae018

	*/

	//数组的 声明方式

	// 1.
	var a [5]string = [5]string{"34124","!@34134","@#532","@3542","@35"}

	//2.
	b:= [5]string {"34124","!@34134","@#532","@3542","@35"}


	fmt.Printf("a = %v\n",a)
	fmt.Printf("b = %v\n",b)


	// 使用for -range遍历菽粟

	for i,v := range(b){
		fmt.Printf("a = %v,b = %v\n",i,v)
	}

	// 1.注意：数组声明了，长度不可变
	/*2. 类型不能变化
	  3、长度固定，不能变化，否则月结
	  4、默认值是类型的默认值
	  5、不能月结，要不就得用panic
	  6、默认值拷贝，所以在函数中对数组进行改变，不会对外面的数组产生影响！！
	  7、跟c不一样，数组的首地址不是数组名

	*/


	// 练习

	// printtt() 

	randomNum()

}

func printtt(){
	var arr [26]byte

	for i:= 0;i<len(arr);i++{
		arr[i]= byte(97+i)
		fmt.Printf(" %c\t",arr[i])

	}
}

func randomNum(){
	// 首先生成一个随机的数组
	arr := [5]int{} 
	rand.Seed(time.Now().Unix())

	for i := 0;i<len(arr);i++{
		arr[i] = rand.Intn(100)
	}
	fmt.Println("arr = ",arr)

	// 然后对他进行倒置
	fmt.Println("len arr = ",len(arr))

	for i := 0;i<len(arr)/2;i++{
		tem := arr[len(arr)-i-1]
		arr[len(arr)-i-1] = arr[i]
		arr[i] = tem
	}
	fmt.Println("arr = ",arr)

}
