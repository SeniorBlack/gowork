package main

import (
	"fmt"
	"reflect"
)

func main() {
	//切片
	//Go语言中，切片是长度可变、容量固定的相同的元素序列。
	arr := [5]int{1, 2, 3, 4} //数组
	//切片定义方法
	//1.通过数组定义
	myslice := arr[0:2] //切片
	fmt.Println(myslice)
	myslice[0] = 9
	qiepian := []int{1, 2} //切片
	fmt.Println(reflect.TypeOf(myslice))
	fmt.Println(reflect.TypeOf(arr))
	fmt.Println(reflect.TypeOf(qiepian))
	/*切片的创建有4种方式：
	1）make ( []Type ,length, capacity )
	2) make ( []Type, length)
	3) []Type{}
	4) []Type{value1 , value2 , ... , valueN }
	*/
	youslice := make([]int, 5, 10)
	fmt.Println(youslice)
	fmt.Println(cap(youslice)) //计算切片容量
	fmt.Println(len(youslice)) //计算切片长度
	//内建函数append()位切片增加元素
	youslice = append(youslice, 6)
	fmt.Println(youslice)
	//可以通过赋值给 _ 来忽略序号和值。

	//如果只需要索引值，去掉 “ , value ” 的部分即可。
	//切片持有对底层数组的引用，如果你将一个切片赋值给另一个，二者都将引用同一个数组。
	//如果函数接受一个切片作为参数，那么其对切片的元素所做的改动，对于调用者是可见的，好比是传递了一个底层数组的指针。

	//在go语言中，遍历切片或数组有两种方式
	//1.传统方式
	msslice := []int{1, 2, 5, 6, 7}
	for i := 0; i < len(msslice); i++ {
		fmt.Println("msslice[", i, "] =", msslice[i])
	}
	//2.go风格遍历,range表达式有两个返回值，第一个是索引，第二个是元素的值：
	for i, v := range msslice {
		fmt.Println("msslice[", i, "] =", v)
	}
	/*动态增减元素是数组切片比数组更为强大的功能。与数组相比，数组切片多了一个存储能
	力（capacity）的概念，即元素个数和分配的空间可以是两个不同的值。合理地设置存储能力的
	值，可以大幅降低数组切片内部重新分配内存和搬送内存块的频率，从而大大提高程序性能。
	假如你明确知道当前创建的数组切片最多可能需要存储的元素个数为50，那么如果你设置的
	存储能力小于50，比如20，那么在元素超过20时，底层将会发生至少一次这样的动作——重新分
	配一块“够大”的内存，并且需要把内容从原来的内存块复制到新分配的内存块，这会产生比较
	明显的开销。给“够大”这两个字加上引号的原因是系统并不知道多大才是够大，所以只是一个
	简单的猜测。比如，将原有的内存空间扩大两倍，但两倍并不一定够，所以之前提到的内存重新
	分配和内容复制的过程很有可能发生多次，从而明显降低系统的整体性能。但如果你知道最大是
	50并且一开始就设置存储能力为50，那么之后就不会发生这样非常耗费CPU的动作，从而达到空
	间换时间的效果。
	数组切片支持Go语言内置的cap()函数和len()函数，代码清单2-2简单示范了这两个内置
	函数的用法。可以看出， cap()函数返回的是数组切片分配的空间大小，而len()函数返回的是
	数组切片中当前所存储的元素个数。*/
}
