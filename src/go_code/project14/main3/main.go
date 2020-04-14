
package main
//golang中的常用函数

import(
	"fmt"
	"strconv"
	"strings"
)

func main(){


	// 字符串常用系统函数

	// 1.len(str) 获得 长度 ,按字节返回，因为一个汉字占三个字节，所以输出为9
	var a string ="123你abc"
	fmt.Printf("a.len = %d\n",len(a))
	// 2.处理有中文的字符串，得把他转成切片 []rune
	str := []rune(a)
	for i:= 0;i<len(str);i++{
		fmt.Printf("a[%d] = %q\n",i,str[i])
	}


	//3. 把字符串转成int ，用strconv下的atoi()
	// 只能检测数字的字符串，如果不是就报错
	var s string = "123456"
	// Atoi返回两个值：int 和 错误
	num , err := strconv.Atoi(s)
	if err != nil{
		fmt.Println("错误为：",err)
	}else {
		fmt.Printf("转换的结果为%d",num)
	}
	// 4. 把int 转成string  strconv下的itoa()

	fmt.Println("\n转成string为：",strconv.Itoa(54321))

	// 5.byote转string 

	sttt := string([]byte{97,98,99})
	fmt.Println("byte{97,98,99}转string为:",sttt)

	// 6. 十进制转其他进制   strconv 中的FormatInt（数，进制）

	fmt.Println("123 十进制转2进制为",strconv.FormatInt(123,2))
	fmt.Println("123 十进制转8进制为",strconv.FormatInt(123,8))
	fmt.Println("123 十进制转16进制为",strconv.FormatInt(123,16))


	// 7.判断字符串是否包含某子串  strings中的Contains
	// abcdefg中是否有bcd
	var flag bool = strings.Contains("abcdefg","bcd")
	if flag {
		fmt.Println("abcdefg中有bcd！！")
	}else {
		fmt.Println("abcdefg中没有bcd！！")
	}

	// 8.判断字符串里有几个子串
	nums := strings.Count("ababababab","ab")

	fmt.Printf("\nababababab中有%d个ab",nums)

	// 9.不区分大小写的字母比较 strings.EqualFold，== 区分大小写

	bbbb := strings.EqualFold("abc","Abc")

	fmt.Printf("\nabc 和 Abc为：%v",bbbb)

	fmt.Printf("\nabc == Abc ? ：%v","abc" == "Abc")


	// 10.在字符串里找子串第一次出现的位置 strings.Index(S,s) ,扎到返回索引，找不到返回-1

	index := strings.Index("12345abc","abc")

	fmt.Printf("\nabc在12345abc 的第 %v 个位置",index)

	// 11.在字符串里找子串最后一次出现的位置 strings.LastIndex(S,s) ,扎到返回索引，找不到返回-1


	// 12.在字符串里找特定子串出现的位置，然后置换成其他子串 strings.Replace(S,s旧，s新，n) ,-1代表全部替换，写几就代表替换几个



	str3 := "go go go jxd"
	//这句话就代表，把str3里的前两个go换成hello
	str4 := strings.Replace(str3,"go","hello",2)
	// 并不会改变原来的子串，replace会返回一个新的子串
	fmt.Println("\n替换之前为：",str3)
	fmt.Println("替换之后为：",str4)


	// 13.将一个字符串拆分成数组  strings.split(第一个参数是字符串，第二个参数是用什么拆分)

	// 我这里用空格去拆str3,很replace一样，这是值拷贝，并不会改变原有字符串
	strArr := strings.Split(str3," ")

	fmt.Println("str3拆分之后，strArr = ",strArr)
	for i:= 0;i<len(strArr);i++{
		fmt.Printf("strArr[%d] = %s\n",i,strArr[i])
	}

	//14.对字符串进行大小写的转换 strings.ToLower,strings.ToUpper

	str6 := "abc"
	fmt.Println("abc变大写之后为",strings.ToUpper(str6))
	//abc变大写之后为 ABC
	str7 := "ABC"
	fmt.Println("ABC变Xiao写之后为",strings.ToLower(str7))
	//abc变小写之后为 abc


	// 15.把字符串两变的空格去掉  strings.TrimSpace()

	str7 = "   hello    !   "

	fmt.Printf("%q两边的空格去掉之后为%q\n",str7,strings.TrimSpace(str7))

	//16.把字符串两边指定的字符去掉  strings.Trim（）

	str7 = "!  hello   !"
	fmt.Printf("%q两边的 ! 去掉之后为%q\n",str7,strings.Trim(str7,"!"))

	//17.把字符串左右的指定的字符去掉  strings.TrimLeft（）  strings.TrimRight（）

	fmt.Printf("%q左边的 ! 去掉之后为%q\n",str7,strings.TrimLeft(str7,"!"))
	fmt.Printf("%q右边的 ! 去掉之后为%q\n",str7,strings.TrimRight(str7,"!"))

	// 18.判断字符串是否以指定的子串开头，url

	str7 = "https://www.baidu.com"
	fmt.Printf("%q是以https开头的吗：%v\n",str7,strings.HasPrefix(str7,"https"))
	// 19.判断字符串是否以指定的子串结束，图片视频

	str7 = "010101.webp"
	fmt.Printf("%q是以webp结尾的吗：%v\n",str7,strings.HasSuffix(str7,"webp"))







}
