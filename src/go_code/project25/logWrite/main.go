package main


import (
	"fmt"
	"os"
	"bufio"
	"io/ioutil"

	
)
func write1(path string){
	// openfile函数能够以指定模式打开文件，从而实现文件写入相关功能。
		// func OpenFile(name string, flag int, perm FileMode) (*File, error) {
		// 	...
		// }
	// 1.name：要打开的文件名 
	// 2.flag：打开文件的模式。 模式有以下几种：
	// 模式	含义
		// os.O_WRONLY	只写
		// os.O_CREATE	创建文件
		// os.O_RDONLY	只读
		// os.O_RDWR	读写
		// os.O_TRUNC	清空
		// os.O_APPEND	追加
	// 3.perm：文件权限，一个八进制数。r（读）04，w（写）02，x（执行）01。
	file ,err :=os.OpenFile(path,os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0666)

	if err != nil{
		fmt.Println("write file failed!err =",err)
		return 
	}
	defer file.Close()
	str := "hello!"
	// 需要转为byte切片
	file.Write([]byte(str))
	// 直接写入string
	file.WriteString(str)
}

/*bufio的write方法*/
func bufWrite2(path string){
	file ,err :=os.OpenFile(path,os.O_RDWR|os.O_APPEND,0666)
	if err != nil{
		fmt.Println("write file failed ,err =",err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i:= 0;i<10;i++{
		//将数据先写入缓存
		writer.WriteString("hwllo122222\n")
	}
	// //将缓存中的内容写入文件
	writer.Flush()


}

/*第三种方式，用ioutil*/
func ioutilWrite(path string,str string){
	err := ioutil.WriteFile(path,[]byte(str), 0666)
	if  err != nil{
		fmt.Println("ioutil write file failed ,err =",err)
	}
}

func main(){

	var path string = "./hello.txt"

	// 1.第一种方式，write和writestring
	write1(path)
	// 2.第二种方式	 bufio
	bufWrite2(path)
	// 3.第三种方式 ，ioutil
	ioutilWrite(path,"我是ioutil写入的")



}