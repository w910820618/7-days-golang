package main

import (
	"log"
	"runtime"
)

var intMap map[int]int

func main() {

	printMemStats("初始化")

	// 添加1w个map值
	intMap = make(map[int]int, 10000)
	for i := 0; i < 10000; i++ {
		intMap[i] = i
	}

	// 手动进行gc操作
	runtime.GC()
	// 再次查看数据
	printMemStats("增加map数据后")

	log.Println("删除前数组长度：", len(intMap))
	for i := 0; i < 10000; i++ {
		delete(intMap, i)
	}
	log.Println("删除后数组长度：", len(intMap))

	// 再次进行手动GC回收
	runtime.GC()
	printMemStats("删除map数据后")

	// 设置为nil进行回收
	intMap = nil
	runtime.GC()
	printMemStats("设置为nil后")

}

func printMemStats(mag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("%v：分配的内存 = %vKB, GC的次数 = %v\n", mag, m.Alloc/1024, m.NumGC)
}
