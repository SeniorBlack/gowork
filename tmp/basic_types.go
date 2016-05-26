package main

import "fmt"

func main() {
	fmt.Println("Hello, Golang!")
	//基本类型(basic types)
	/*整型:int int8 int16 int32 int64;
	uint uint8 uint16 uint32 uint64 uintptr
	byte //uint8 的别名
	rune //int32的别名，代表一个Unicode码
	int，uint 和 uintptr 类型在32位的系统上一般是32位，
	而在64位系统上是64位。当你需要使用一个整数类型时，你
	应该首选 int，仅当有特别的理由才使用定长整数类型或者无符号整数类型。
	*/
	var a int = 100
	var b byte = 12
	fmt.Printf("整数：\na = %d\nb = %d\n", a, b)

	//bool
	//true false
	c := true
	fmt.Println("bool:")
	fmt.Println(c)

	//浮点型
	//float32 float64
	var d float32
	d = 1.226
	fmt.Printf("浮点数：\nc = %f\n", d)

	//复数
	//complex64 complex128
	var e complex64 = 2.3 + 3i
	fmt.Println("复数:")
	fmt.Println(e)

	//字符串
	var f string = "Hello, golang!"
	fmt.Println(f)

	//error类型

	/*变量在定义时没有明确的初始化会赋值为 “零值”
	零值是：
		数值类型为 0，
		布尔类型为 false,
		字符串为 “” （空字符串）
	*/

	/*
		类型转换：
		表达式T(v),将值v转换为类型T
		var i int = 42
		var f float64 = float64(i)
		与 C 不同的是 Go 的在不同类型之间的项目赋值时需要显式转换
	*/

	//常亮定义用const
	const world = "world"
	//常量可以是字符，字符串，布尔或数字类型的值
	//常量不能使用:=语法定义

}
