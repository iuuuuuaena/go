package main

// string和slice

import (
	"fmt"
)

func main() {
	// string 底层是个byte，所以他底层是个切片

	// string可以进行切片处理
	var str string = "abcdefghig"

	// 切片第六个位置到最后
	slice := str[6:]

	fmt.Println(slice)

	// string不可变
	// 要想改变string，先把他转成byte切片

	arr := []byte(str)

	fmt.Println(arr)
	// 这时候，输出就是abcdefgh对应的acill值
	// [97 98 99 100 101 102 103 104 105 103]

	arr[4] = 69
	// 修改第四个位置的值
	arr[5] = 'A'

	fmt.Println(arr)
	// 可以看到改变了
	// [97 98 99 100 1 102 103 104 105 103]

	str = string(arr)

	fmt.Println(str)
	//abcdEAghig
	// 这样就可以修改string的值了！！1
	// 这种方法的缺点，就是不能用汉字，因为byte是按照字节存储的，而汉字占三个字节，
	// 用byte存储就会把汉字拆分，编程乱码

	// 所以我们推荐用切片去分解。

	string2 := "你好啊！hello！"
	slice2 := []rune(string2)
	fmt.Printf("%s\n", string2)
	fmt.Printf("%c\n", slice2)
	// 你好啊！hello！

	fmt.Printf("%c\n", slice2[1])
	fmt.Printf("%c\n", slice2[1])
	slice2[1] = '啊'
	string2 = string(slice2)

	fmt.Printf("%s\n", string2)
	// 你啊啊！hello！

}
