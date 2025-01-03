package main

import (
	"fmt"
	"sort"
)

func main() {
	// 定义一个整数切片
	numbers := []int{5, 3, 7, 1, 9, 2}

	// 打印排序前的切片
	fmt.Println("Before sorting:", numbers)

	// 使用 sort.Ints 对切片进行排序
	sort.Ints(numbers)

	// 打印排序后的切片
	fmt.Println("After sorting:", numbers)

	// 检查是否已经排好序
	if sort.IntsAreSorted(numbers) {
		fmt.Println("The slice is sorted.")
	} else {
		fmt.Println("The slice is not sorted.")
	}
}
