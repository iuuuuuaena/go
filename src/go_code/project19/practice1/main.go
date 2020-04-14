//小练习
package main

import (
	"fmt"
	"time"
)

/**

编写一个函数，实现可以让用户登录，修改密码的小程序
5
*/

// 定义存放用户数据的结构体
type Users struct {
	U string
	P string
}

var userList map[string]Users

var count *int = new(int)

// 定义map切片，并初始化
var userMap []map[string]Users

func Init() {
	userMap = make([]map[string]Users, 10)
	// userList = make (map[string]Users,2)
	*count = 0
}

// 向map切片中插入map
func add(mapSlice []map[string]Users, user Users, count *int) {
	i := fmt.Sprintf("%d", *count)
	if mapSlice[0] == nil {
		mapSlice[0] = make(map[string]Users, 2)

		mapSlice[0]["U"+i] = user
	}
	mapSlice[0]["use"+i] = user
	fmt.Println("注册成功")
}

// 查询map切片中是否账号密码一致
func search(mapSlice []map[string]Users, user Users) {
	var flag bool = false
	for _, V := range mapSlice {
		for k, v := range V {
			if user.U == v.U && user.P == v.P {
				fmt.Printf("\n登录成功，您现在的账号顺序是%s\n", k)
				flag = true
				return
			} else {
				continue
			}
		}
	}
	if flag == false {
		fmt.Println("\n登录失败，账号密码错误！")
		return
	}

}

func stop() {
	time.Sleep(time.Millisecond * 1000)
}
func changePass(mapSlice []map[string]Users, user Users) {

	for k, v := range mapSlice[0] {
		if user.U == v.U && user.P == v.P {
			fmt.Println("\n账号密码正确，请输入新密码！\n", k)
			nPass := ""
			fmt.Scanf("%s", &nPass)
			v.P = nPass
			fmt.Println("稍等！")
			stop()
			fmt.Println("\n密码修改成功！\n", k)
			break
			return
		}
	}

}

func show() {
	fmt.Println("请输入编号：")
	fmt.Println("1.log in")
	fmt.Println("2.log up")
	fmt.Println("3.change password")
	fmt.Println("4.show all members")
	fmt.Println("5.quit")
}

func display(mapSlice []map[string]Users) {
	fmt.Println("所有用户信息为：")
	for _, V := range mapSlice {
		for k, v := range V {
			fmt.Printf("%s:\n\tname = %s\tpass = %s\n", k, v.U, v.P)
		}
	}

}

func main() {

	fmt.Println("欢迎登录！！！")

	//初始化
	Init()
	var flag bool = false
	for {
		show()
		var in int
		fmt.Scanf("%d", &in)
		switch in {
		case 1:
			fmt.Println("稍等！")
			stop()
			fmt.Println("请输入你的账号和密码！")
			var username string
			var password string
			fmt.Scanf("%s %s", &username, &password)
			fmt.Println(username, password)
			user1 := Users{username, password}
			search(userMap, user1)
		case 2:

			fmt.Println("稍等！")
			stop()
			fmt.Println("请输入你的账号和密码！")
			var username string
			var password string
			fmt.Scanf("%s %s", &username, &password)
			fmt.Println(username, password)

			var user1 Users
			user1 = Users{username, password}
			number := ""
			number = fmt.Sprintf("%d", *count)
			fmt.Println("count = ", number)

			add(userMap, user1, count)
			*count++
		case 3:

			fmt.Println("稍等！")
			stop()
			fmt.Println("请输入你的账号和密码！")
			var username string
			var password string
			fmt.Scanf("%s %s", &username, &password)
			user1 := Users{username, password}
			changePass(userMap, user1)

		case 4:
			fmt.Println("稍等！")
			stop()
			display(userMap)

		case 5:
			fmt.Println("稍等！")
			stop()
			flag = true
		default:
			fmt.Println("请重新选择！！1")

		}
		if flag == true {
			break
		}

	}

}
