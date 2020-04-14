//指针


package point

import (
	"fmt"
)
/*
	练习指针的使用
*/
func main(){

	//取地址  &
	//取值 *

	// 定义int类型的a
	a := 10

	fmt.Printf("a = %d,a的默认类型是%T,a的地址为%p\n",a,a,&a)
	//a = 0,a的地址为0xc0000b4008%


	
	var p  *int  = &a

	fmt.Printf("p = %p,p的地址为%p,p保存的地址保存的值为%d\n",p,&p,*p)

	
	fmt.Printf("a = *p？%v \n&a == p? %v\n ",a == *p,&a == p)

	//可以通过p来修改a的值

	*p = 100
	fmt.Printf("a = %d,a的默认类型是%T,a的地址为%p\n",a,a,&a)

	fmt.Printf("p = %p,p的地址为%p,p保存的地址保存的值为%d\n",p,&p,*p)

	/* 输出结果
	a = 10,a的默认类型是int,a的地址为0xc0000b4008
	p = 0xc0000b4008,p的地址为0xc0000ae020,p保存的地址保存的值为10
	a = *p？true 
	&a == p? true
	a = 100,a的默认类型是int,a的地址为0xc0000b4008
	p = 0xc0000b4008,p的地址为0xc0000ae020,p保存的地址保存的值为100
*/

	



}