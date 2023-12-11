package main

import "fmt"

/*
*
快速排序
*/
func main() {
	arr := []int{1, 5, 3, 33, 22, 11}
	fmt.Printf("befor:%v\n", arr)
	descArr := quickSort(arr,0, len(arr)-1)
	fmt.Printf("after:%v", descArr)
}

/**
从小到大快速排序
时间复杂度： O(n^2)
空间复杂读： O(1)
 */
func quickSort(arr []int,start int,end int) []int {

if(start<end){

}
	return arr
}



//func quickSort(arr []int, low, high int) {
//	if low < high {
//		pivotIndex := partition(arr, low, high)
//
//		quickSort(arr, low, pivotIndex-1)
//		quickSort(arr, pivotIndex+1, high)
//	}
//}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
