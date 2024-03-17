package reviewGo

import (
	"fmt"
	"sort"
)

func RunDay24() {
	day24Ex3()
}

func day24Ex1() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println(arr[3])
}

func day24Ex2() {
	sli := make([]int, 10)
	fmt.Println(sli)
	fruitSli := []string{"apple", "banana", "orange"}

	fruitSli = append(fruitSli, "durian", "CHerryTOmato", "Blueberry")
	fmt.Println(fruitSli)
	newFru := fruitSli[:3]
	fmt.Println(newFru)

	fruitSli = append(fruitSli[:2], fruitSli[4:]...)
	for i, v := range fruitSli {
		fmt.Println(i, v)
	}
}

func day24Ex3() {
	fruitInfo := make(map[string]int)
	fruitInfo["apple"] = 1
	fruitInfo["banana"] = 2
	fruitInfo["orange"] = 3
	fruitInfo["cheery"] = 4
	fmt.Println(fruitInfo)

	keyArr := make([]string, 0, len(fruitInfo))
	for k, v := range fruitInfo {
		fmt.Println(k, v)
		keyArr = append(keyArr, k)
	}
	sort.Strings(keyArr)
	for _, i := range keyArr {
		fmt.Println(i, fruitInfo[i])
	}
	_, elementExists := fruitInfo["apple"]
	fmt.Println(elementExists)
}
