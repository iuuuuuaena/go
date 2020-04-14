package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"io"
	"bufio"
)
//os库用于文件操作
//bufio 缓存区读取




/*for循环文件读取操作*/
func readFile(path string){
	fileObj,error := os.Open(path)

	if error != nil{
		fmt.Println("open file failed!")
		return 
	}
	defer fileObj.Close()
	var by [128]byte
	for {
		// by = make([]byte,128)
		n,err := fileObj.Read(by[:])
		// EOF是io；里的标识达到文件结尾的意思，就是说，已经读到文件结尾了，就结束
		if err == io.EOF{
			return
		}
		if err!=nil{
			fmt.Println("read failed!")
			return 
		}
		fmt.Printf("读了%d个字节\n",n)
		fmt.Println(string(by[:n]))
	}

}
// 日志读取问题！！！11


func bufRead(path string){

	file ,err := os.Open(path)
	if err != nil{
		fmt.Println("文件读取失败！")
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for{
		line,err := reader.ReadString('\n')
		if err == io.EOF{
			if len(line) != 0{
				fmt.Println(line)
			} 
			fmt.Println("读完了！")
			break
		}
		if err != nil{
			fmt.Println("read file failed ,err = ",err)
		}
		fmt.Print(line)
	}

}

/*使用ioutil读取content*/
func ioutilRead(path string){
	content,err :=  ioutil.ReadFile(path)
	if err != nil{
		fmt.Println("read file failed!,err =",err)
		return 
	}
	fmt.Println(string(content))

}


func main(){

	path := "./main.go"
	// 1.for循环读取
	readFile(path)
	fmt.Println("------------------------")
	// 2,用bufio读取
	bufRead(path)
	// 3.还有最简单的方式，用ioutil
	ioutilRead(path)

}