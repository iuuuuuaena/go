package main


import (
	"fmt"
	"io"
	"os"
)


/*copy一个文件到另一个文件*/
func copy(path1 string,path2 string)(written int64,err error){
	resource ,err1 := os.Open(path1)
	if err1 != nil{
		fmt.Println("file1 read failed ,err= ",err1)
		return 
	}
	defer resource.Close()


	distination ,err2 := os.OpenFile(path2,os.O_CREATE|os.O_RDWR|os.O_TRUNC,0666)
	if err2 != nil{
		fmt.Println("file2 write failed ,err =",err2)
		return 
	}
	defer distination.Close()

	// 直接返回copy之后的
	return io.Copy(resource,distination)

}

func main(){

	path1 := "./path1.txt"
	path2 := "./path2.txt"
	_,err := copy(path1,path2)
	if err != nil{
		fmt.Println("copy failed!err = ",err)
		return
	}
	fmt.Println("copy success!!!!1")

}