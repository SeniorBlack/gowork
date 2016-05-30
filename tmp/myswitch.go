package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go run on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Println(os, ".")
	}
	/*
		在使用switch结构时，我们需要注意以下几点：
		左花括号{必须与switch处于同一行；
		条件表达式不限制为常量或者整数；
		单个case中，可以出现多个结果选项；
		与C语言等规则相反， Go语言不需要用break来明确退出一个case；
		只有在case中明确添加fallthrough关键字，才会继续执行紧跟的下一个case；
		可以 不设定switch之 后的条 件表达式， 在此种情况 下，整个switch结构与 多个
		if...else...的逻辑作用等同。
	*/
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today)
	fmt.Println(reflect.TypeOf(today))
	fmt.Println(time.Saturday)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
