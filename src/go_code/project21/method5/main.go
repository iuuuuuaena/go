// 复习下结构体的创建

package main


import "fmt"

type Student struct{
	Name string
	Age int
}


func main(){


	// 1.
	var student1 Student = Student{"tom",1}
	// 2
	student2 := Student{"jarry",2}

	var student3 = Student{"mike",3}

	var student4 = Student{
		Name:"sarry",
		Age:123,
	}

	fmt.Println(student1,student2,student3,student4)
	// {tom 1} {jarry 2} {mike 3} {sarry 123}   都没问题



	// 4.返回结构体指针 ，对象指向一个地址，


	var student5 = &Student{"iii",5}


	var student6 = &Student{
		Name :"jjj",
		Age :6,
	}

	student7 := &Student{"yyy",7}


	fmt.Println(*student5,*student6,*student7)
	// {iii 5} {jjj 6} {yyy 7}






}