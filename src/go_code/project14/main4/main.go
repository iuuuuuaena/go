package main

import(
	"fmt"
	"time"
	"strconv"

)
// 常用的时间函数

func main(){

	//time.Time是一种数据类型，专门表示时间

	//1.获取当前时间

	now := time.Now()

	fmt.Printf("now是%T类型的，当前时间为：%v",now,now)

	// 数据结果为：now是time.Time类型的，当前时间为：2020-03-19 17:28:05.089156 +0800 CST m=+0.000084747%                                            

	// 使用time的方式获取年月日

	fmt.Printf("\n年为：%v\n月为：%v\n日为：%v\n当前时间为:%v:%v:%v",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())

	// 简写一下为：
	// 注意：换行的时候，要把, 留在上面
	fmt.Printf("\n当前时间为: %02d-%02d-%02d %02d:%02d:%02d",now.Year(),now.Month(),
	now.Day(),now.Hour(),now.Minute(),now.Second())

	// 2.格式话日期的第二种方式,这种方式是固定的，必须是2006年1月2日 。。。。。
	fmt.Println()
	fmt.Println(now.Format("2006-01-02 15:04:05"))


	//3.时间常量

	/*
	秒     time.Second 
	毫秒   time.Millisecond 
	微妙   time.Microsecond 
	纳秒   time.Nanosecond 

	*/


	// 编写一个函数，隔一段时间输出一个

	for i:=0 ;i<10;i++{
		fmt.Printf("\n%d\n",i)
		// 间隔100毫秒输出一次
		time.Sleep(time.Millisecond * 100)
		
	}
	i:=0
	for{
		i++
		fmt.Println(i)
		time.Sleep(time.Millisecond*10)
		if i >= 10{
			break
		}
	}

	// 4.获取随机数字，用UNix和，unixNano


	start := time.Now().Unix()

	calculateTime()

	end := time.Now().Unix()

	fmt.Printf("\n执行这个函数用了%v喵",end-start)




}

func calculateTime(){
	str := ""
	for i:=0;i<100;i++{
		time.Sleep(time.Millisecond* 100)
		fmt.Println(str+" "+strconv.Itoa(i))
	}
}
