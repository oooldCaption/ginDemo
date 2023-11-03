package reviewGo

import (
	"fmt"
)

func RunDay17() {
	runDay17Ex1()
}

func runDay17Ex1() {
	var tMap = map[string]string{}
	tMap["name"] = "dipsy"
	tMap["age"] = "30"
	tMap["gener"] = "男"
	fmt.Println(tMap)
	var map1 = map[string][]int{}
	map1["name"] = []int{1, 2, 3, 4, 5}
	map1["age"] = []int{1, 2, 3, 4, 5}
	fmt.Println(map1)
}
