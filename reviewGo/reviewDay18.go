package reviewGo

import (
	"fmt"
	"ginDemo/reviewGo/model"
)

func RunDay18() {
	day18Ex1()
}
func day18Ex1() {
	cat := model.NewSmallCat("小花花", 23, 8.2)
	fmt.Println(cat)
	cat.SetAge(4)
	fmt.Println(cat)
}

func findCat() {
	cat1 := map[string]string{"name": "小白", "age": "5"}
	cat2 := map[string]string{"name": "小灰", "age": "3"}

	carArr := [2]map[string]string{cat1, cat2}
	name := "小白"
	for _, m := range carArr {
		if m["name"] == name {
			print(m["age"])
		}
	}
}
func findCatWithStruct() {
	type Cat struct {
		name  string
		age   int
		color string
	}
	cat1 := Cat{
		name:  "小花花",
		age:   20,
		color: "黑白色",
	}
	cat2 := Cat{
		name:  "小灰",
		age:   22,
		color: "灰色",
	}
	fmt.Println(cat1)
	fmt.Println(cat2)

}

func day18Ex0() {
	type Student struct {
		Name   string
		AgePtr *int
		Class  int
	}
	//s1 := Student{
	//	Name:   "sony",
	//	AgePtr: nil,
	//	Class:  8,
	//}
	//age := 28
	//s1.AgePtr = &age
	//fmt.Println(*s1.AgePtr)
	//age = 39
	//fmt.Println(*s1.AgePtr)

	s2 := &Student{
		Name:   "tt",
		AgePtr: nil,
		Class:  22,
	}
	fmt.Println(s2)

}
