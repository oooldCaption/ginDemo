package alg

import "fmt"

// 是对 插入算法 的一种优化, 利用对问题的而分化,实现递归完成快速排序.  0(nlgn)

/*
1. 将数据根据一个值根据一个值 按照大小分成左右两部分, 左边小于 这个值,  右边大于这个值
2.将两边数据进行递归 然后继续调用 1
3. 将数据合并
*/
func QuickSort(arr []int) []int {
	// 判断数组长度
	if len(arr) <= 1 {
		return arr
	}

	// 取出一个数据
	splitData := arr[0]

	// 比我小的数据
	low := make([]int, 0, 0)
	// 比我大的数据
	height := make([]int, 0, 0)
	// 跟我一样大的数据
	mid := make([]int, 0, 0)
	
	//加入一个数据
	mid = append(mid, splitData)

	for i := 1; i < len(arr); i++ {
		if arr[i] < splitData {
			low = append(low, arr[i])
		} else if arr[i] > splitData {
			height = append(height, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}

	//fmt.Println(low, height, mid)
	low, height = QuickSort(low), QuickSort(height)
	myarr := append(append(low, mid...), height...)
	fmt.Println(myarr)
	return myarr
}
