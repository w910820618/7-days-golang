package main

import (
	"7-days-golang/two-project/lib"
	"fmt"
)

func main() {
	// 常量
	//const str1 string = "hello world"
	//fmt.Println(str1)
	//// 尝试取消注释，看一下效果
	////str1 = str1 + "hello"
	//
	//// 变量
	//var str2 string = "hello world"
	//fmt.Println(str2)
	//
	//str3 := "hello world"
	//fmt.Println(str3)
	//
	//var str4 string
	//str4 = "hello world"
	//fmt.Println(str4)

	// 整型
	// 无符号数
	//var num1 uint
	//var num2 uint8
	//var num3 uint16
	//var num4 uint32
	//var num5 uint64

	// 有符号数
	//var num1 int
	//var num2 int8
	//var num3 int16
	//var num4 int32
	//var num5 int64

	// 浮点型
	// var num1 float32
	// var num2 float64

	// 字符串
	var str = "C语言中文网\nGo语言教程"
	// 按Unicode字符遍历字符串
	for _, s := range str {
		fmt.Printf("Unicode: %c  %d\n", s, s)
	}
	// 按ascii字符遍历字符串
	theme := "start"
	for i := 0; i < len(theme); i++ {
		fmt.Printf("ascii: %c  %d\n", theme[i], theme[i])
	}

	//fmt.Println(str)
	//
	//// 字符型
	//var ch byte = 'A'
	//fmt.Printf("%d - %c - %X - %U\n", ch, ch, ch, ch)
	//
	//// 布尔类型
	var result bool = true
	fmt.Println(result)
	//
	//// 数组与切片
	//var s1 = make([]int, 5)
	//var s2 = make([]int, 5, 10)
	//var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s1 = a[:4]
	//s2 = a[3:7]
	//fmt.Println(s1)
	//fmt.Println(s2)
	//s1[3] = 100
	//fmt.Println(s1)
	//fmt.Println(s2)
	//
	//// 包的引用
	lib.Run()
}
