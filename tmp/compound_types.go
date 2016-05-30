package main

import (
	"fmt"
)

type user struct {
	name string
	age  uint
}

func main() {
	//复合类型(compound types)
	//指针
	//类型 *T 是指向类型 T 的值的指针。其零值是 nil 。
	var p *int
	var a int = 5
	p = &a
	fmt.Println(*p)
	//与c语言不通，go没有指针运算

	//结构体（struct）
	//一个结构体（struct）就是一个字段的集合
	//type可以用来定义一个结构体类型
	/*定义格式:
	type 结构体名称 struct {
		属性1 类型
		属性2 类型
		...
	}
	*/
	//结构体初始化可以使用new关键字和var关键字，不通的是new返回类型的一个指针，
	//使用var,返回结构体本身
	user1 := new(user)
	user1.name = "Chen Haibo"
	user1.age = 35
	fmt.Println(*user1)

	var user2 user
	user2.name = "Wan Zhaoying"
	user2.age = 29
	fmt.Println(user2)

	//go语言特有方法
	user3 := user{"Ma Shuang", 25} //按顺序给属性赋值，
	//user3 := user{age: 25, name: "Ma Shuang"} //指明属性赋值
	fmt.Println(user3)

	//数组
	/*数组定义方法
	[长度]类型{初始值}
	*/
	//其中长度可以者换成三个点“...”，如果这么做的话，系统会根据初始化的值自动定义数组的长度
	//如果指明了长度，但是没有初始值，则会根据数组类型初始化每个值
	arr := [2]int{3, 2}
	//arr := []int{3, 2} //不指定长度代表切片
	//arr := [...]int{3, 2} //...代表数组
	fmt.Println(arr) //输出[3 2]
	// 在c语言中，数组是值类型，将一个数组赋值给另一个数组是一个拷贝过程
	arr2 := arr
	fmt.Println(cap(arr2)) //cap()最大元素个数
	fmt.Println(len(arr2)) //len()实际元素个数
	//数组的cap和len相等

}
