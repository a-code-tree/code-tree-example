package main

import "fmt"

/*
*
冒泡排序
*/
func main() {
	arr := []int{1, 5, 3, 33, 222, 11}
	fmt.Printf("befor:%v\n", arr)
	descArr := bubbleSort(arr)
	fmt.Printf("after:%v", descArr)
}

/*
*
从小到大冒泡排序
时间复杂度： O(n^2)
空间复杂度： O(1)
*/
func bubbleSort(arr []int) []int {
	arrLength := len(arr)
	for i := 0; i < arrLength-1; i++ {
		for j := 0; j < arrLength-1-i; j++ {
			if arr[j] > arr[j+1] {
				//如果前一个数字比后一个数字大，进行交互顺序
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
