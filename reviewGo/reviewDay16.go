package reviewGo

import (
	"fmt"
)

func RunDay16() {
	day16Ex1()
}

func day16Ex1() {
	var slice1 []int
	fmt.Println(slice1)
	slice1 = append(slice1, 1, 2, 3, 4, 5, 6, 7, 454, 4, 3, 2, 2)
	fmt.Println(slice1)

	slice2 := make([]int, 5)
	fmt.Println(slice2)

	slice3 := make([]string, 5, 10)
	fmt.Println(slice3)

	arr := [...]int{1, 2, 3, 4, 5}
	slice4 := arr[:]
	fmt.Println(slice4)
	//for i, i2 := range slice1 {
	//	fmt.Println(i)
	//	fmt.Println("=====")
	//	fmt.Println(i2)
	//}

	//var wg sync.WaitGroup
	//for _, item := range slice1 {
	//	wg.Add(1)
	//	go func() {
	//		fmt.Println(item)
	//		wg.Done()
	//	}()
	//}
	//for _, item := range slice1 {
	//	wg.Add(1)
	//	go func(i int) {
	//		fmt.Println(i)
	//		wg.Done()
	//	}(item)
	//}
	//wg.Wait()
	//testString := "thisIsWay"
	//slice5 := testString[1:3]
	//
	//fmt.Println(slice5)

}
