// string转其他类型

package main

import (
	"fmt"
	"strconv"
)

/* string 转其他数据类型的注意事项
比如我们可以吧123这个string类型的数据转化为int
但是我们不能把hello这个字符串转化为int，这样相当于转化失败，
int那个还会是他的原始值，如果没有赋值，就会是0，
其他类型也是一样。

*/
func main() {
	// string转bool
	var a bool

	var str1 string = "true"

	a, _ = strconv.ParseBool(str1)

	fmt.Printf("由str转化为bool的a为%v\n", a)

	// string转int

	var b int64

	var str2 string = "123"
	// 10进制，int64
	b, _ = strconv.ParseInt(str2, 10, 0)

	fmt.Printf("由str转化为int的b为%d\n", b)

	// string转float
	var c float64 = 0.0

	var str3 string = "100.23"

	c, _ = strconv.ParseFloat(str3, 64)

	fmt.Printf("由str转化为float的c为%f\n", c)

	// 举个栗子，把hello转化为bool

	var d bool

	var str4 string = "hello"

	d, _ = strconv.ParseBool(str4)

	fmt.Printf("由str转化为bool的d为%v\n", d)
	//hello转不了bool,所以输出为false
	// 由str转化为bool的d为false

}
