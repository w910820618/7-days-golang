package main

import "fmt"

func main() {
	// 数组
	//var a [5]int
	//var b [5]int

	//a = [5]int{1, 2, 3, 4, 5}
	//b = a
	//a[1] += 1
	//fmt.Println(a)
	//fmt.Println(b)

	// 切片
	c := make([]int, 5)
	d := make([]int, 5)
	c = []int{1, 2, 3, 4, 5}
	d = c
	c[1] += 1
	fmt.Println(c)
	fmt.Println(d)
}
