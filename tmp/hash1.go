package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func main() {
	TestString := "Hi,pandaman!"

	Md5Inst := md5.New()              //获得md5对象
	Md5Inst.Write([]byte(TestString)) //将TestString转化成[]byte后读入内存
	/*
		这里直接对一串字符串计算MD5。其中通过md5.New()初始化一个MD5对象，其实它是一个hash.Hash对象。
		函数原型为 func New() hash.Hash 。该对象实现了hash.Hash的Sum接口：计算出校验和。其函数原型
		为 func Sum(data []byte) [Size]byte 这里的官方Manual对其的描述我感觉有点问题。其官方描述为:
		" Sum returns the MD5 checksum of the data. "通过翻阅源码可以看到他并不是对data进行校验计算，
		而是对hash.Hash对象内部存储的内容进行校验和 计算然后将其追加到data的后面形成一个新的byte切片。因此通常的使用方法就是将data置为nil
		[]byte和string之间可以相互转换
	*/
	Result := Md5Inst.Sum([]byte(nil))      //通过内部函数checkSum计算出其校验和,生成128bits
	fmt.Printf("%x\n\n", Result)            //按十六进制输出
	fmt.Println(hex.EncodeToString(Result)) //转化成16进制字符串

	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(TestString))
	Result = Sha1Inst.Sum([]byte(nil))
	fmt.Printf("%x\n\n", Result)

}
