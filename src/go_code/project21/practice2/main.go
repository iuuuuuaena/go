package main

// 小练习，结构体学生管理系统

/*用到的知识有
	1.结构体
	2.遍历map
	3.json的序列化

*/
import (
	"fmt"
	"os"
)

/*声明一个全局的结构体*/
var stuMgr studentsMgr

/*定义学生结构体*/
type students struct{	
	ID  int64
	Name string
}
/*定学生管理结构体*/
type studentsMgr struct{
	allStudent map[int64]students
}

/*查看所有学生*/
func (s studentsMgr)showAllStudents(){
	for _ , stu := range s.allStudent{
		fmt.Printf("ID : %d \tName : %s\n",stu.ID,stu.Name)
	}

}
/*添加学生*/
func (s studentsMgr)addStudents(){
	var (
		stuID int64
		stuName string
	)

	fmt.Println("请输入ID")
	fmt.Scanln(&stuID)
	fmt.Println("请输入Name")
	fmt.Scanln(&stuName)
	stu := students{
		ID : stuID,
		Name : stuName,
	}

	s.allStudent[stu.ID] = stu



}
/*删除学生*/
func (s studentsMgr)deleteStudents(){
	// 1.获取用户输入
	fmt.Println("请输入要删除的学生的学号：")
	var stuID int64

	fmt.Scanln(&stuID)
	fmt.Println("你输入的是：",stuID)
	// 2.查找是否有这个学生
	_ ,OK := s.allStudent[stuID]
	if !OK{
		delete(s.allStudent,stuID)
	}else {
		fmt.Println("没有这个学生!")
	}
}
/*修改学生*/
func (s studentsMgr)editStudents(){
	// 1.查找是否有这个学生
	fmt.Println("请输入要修改的学生的学号：")
	var stuID int64
	fmt.Scanln(&stuID)
	fmt.Println("你输入的学号为",stuID)

	value ,OK := s.allStudent[stuID]

	if !OK {
		fmt.Println("对不起，查无此人！")
		return
	}
	fmt.Printf("这位学号为 %d的学生，当前的姓名为%v\n",stuID,value)
	fmt.Println("请输入修改后的姓名：")
	var newName string 
	fmt.Scanln(&newName)
	newStu := students{
		ID : stuID,
		Name : newName,
	}
	s.allStudent[stuID] = newStu
	fmt.Printf("这位学号为 %d的学生，修改后的姓名为%v\n",stuID,s.allStudent[stuID].Name)
	



}



func showMess(){
	fmt.Println("请选择操作：")
	fmt.Println(`
		1.展示所有学生
		2.添加学生
		3.编辑学生
		4.删除学生
		5.退出系统
	`)
}


func main(){

	stuMgr =studentsMgr{
		allStudent: make(map[int64]students,10),
	}

	for {
		showMess()

		fmt.Println("请选择：")
		var choice int
		fmt.Scanln(&choice)
		switch choice{
		case 1:
			stuMgr.showAllStudents()
		case 2:
			stuMgr.addStudents()
		case 3:
			stuMgr.editStudents()
		case 4:
			stuMgr.deleteStudents()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("gum!")

		}


	}
}