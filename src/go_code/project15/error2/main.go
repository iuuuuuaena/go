package main

import(
	"fmt"
	"errors"
)

func readConf(name string)(err error){
	if name == "config.ini"{
		return nil
	}else {
		//返回一个自定义错误
		return errors.New("文件名错误！！！")
	}


}

func main(){
	err := readConf("nfig.ini")
	// 传的文件不对，就是输出panic的语句
	if err != nil{
		//如果文件错误，就返回错误并终止程序！！！
		panic(err)
	}
	fmt.Println("------")

	// panic: 文件名错误！！！
	// goroutine 1 [running]:
	// main.main()
	// 		/Users/mac/go/src/go_code/project15/error2/main.go:23 +0x5a
	// exit status 2
}