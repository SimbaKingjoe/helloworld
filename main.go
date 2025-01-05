// main.go
package main

import (
	"fmt"
	"mypackage" // 根据你的项目路径调整这里的导入路径
)

func main() {
	// 整数排序示例
	intSlice := []int{5, 3, 7, 1, 9, 2}
	fmt.Println("Before sorting (ints):", intSlice)
	mypackage.SortInts(intSlice)
	fmt.Println("After sorting (ints):", intSlice)

	// 字符串按长度排序示例
	stringSlice := []string{"apple", "banana", "grape", "peach", "watermelon"}
	fmt.Println("\nBefore sorting by length (strings):", stringSlice)
	mypackage.SortStringsByLength(stringSlice)
	fmt.Println("After sorting by length (strings):", stringSlice)
}
