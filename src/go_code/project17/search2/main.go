package main

import "fmt"

// 二分查找的变种，插值查找
// /先来看一个实际问题：
// 我们在一本英汉字典中寻找单词“worst”，我们决不会仿照对半查找（或Fibonacci查找）那样，先查找字典中间的元素，然后查找字典四

// 分之三处的元素等等. 事实上，我们是在所期望的地址（在字典的很靠后的地方）附近开始查找的，我们称这样的查找为插值查找.
// 可见，插值查找不同于前面讨论的几种查找算法，前面介绍的查找算法是基于严格比较的，即假定我们对线性表中元素的分布一无所知（或

// 称没有启发式信息）. 然而实际中，很多查找问题所涉及的表满足某些统计的特点.

// 插值查找在实际使用时，一般要满足两个假设条件：

// （1）每一次对数据的访问与通常的指令相比，费用都是相当昂贵的。例如，待查找的表一定是在磁盘而非内存中，因而每一次比较都

// 要进行磁盘访问。

// （2）数据不仅是已被排好序的，而且呈现均匀分布特征。

func main(){


	arr := [7]int {1,2,3,4,5,6,7}
	fmt.Println("arr=",arr)
	ret := insertSearch(&arr,0,6,6)
	fmt.Println("ret=",ret)
	
}

func insertSearch(arr *[7]int ,left int ,right int ,target int )(index int ){
	// 判断条件相对于二分查找，多了范围的界定，主要是为了防止mid过域
	if left >right || target < (*arr)[0] || target > (*arr)[len(*arr)-1]{
		return -1
	}
	// 插值算法的核心公式
	mid := left + (right - left)*(target - (*arr)[left])/((*arr)[right] -(*arr)[left])
	if target > (*arr)[mid]{
		insertSearch(arr,mid+1,right,target)
	}else if target < (*arr)[mid]{
		insertSearch(arr,left,mid-1,target)
	}else{
		return mid
	}
	return -1

}