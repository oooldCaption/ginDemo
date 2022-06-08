package main

import (
	"fmt"
	"ginDemo/algorithm/alg"
)

func main() {
	t := []int{1, 2, 3, 22, 121, 223, 32211, 3332, 445, 3231, 2123131, 1123, 44454, 21321321, 2132132}
	//fmt.Println(t)
	arr := alg.QuickSort(t)
	fmt.Println(arr)

}
