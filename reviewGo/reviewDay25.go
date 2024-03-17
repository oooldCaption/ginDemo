package reviewGo

import (
	"fmt"
	"reflect"
)

func RunDay25() {
	//testParams(1, 2, 3, 4, 5, 5)
	testif()
}
func testif() {
	num := 10
	if num >= 5 {
		fmt.Println("1")
	} else if num >= 8 {
		fmt.Println("2")
	}
}

func returnTwoResult(num1, num2 int) (sum, obb int) {
	return num1 + num2, num1 - num2
}

func testParams(nums ...int) {

	for i, v := range nums {
		fmt.Println(i, v)
	}
	fmt.Println(reflect.TypeOf(nums))
}
