package reviewGo

import (
	"fmt"
	"strings"
)

func RunDay15() {
	day15Ex2()
}

func closure() func() string {
	x := 18
	return func() string {
		x += 1
		result := ""
		switch x {
		case 18:
			result = "那年十八岁"
		case 19:
			result = "那年十9岁"
		case 20:
			result = "二十岁啦"
		case 21:
			result = "好好读书"

		}
		return result
	}
}

func hasEndFix(endFix string) func(fileName string) string {
	s := endFix
	return func(fileName string) string {
		if strings.HasSuffix(fileName, s) {
			return fileName
		}
		return fileName + s
	}
}
func day15Ex2() {
	c := closure()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
	f := hasEndFix(".jpg")
	fmt.Println(f("demo"))
	fmt.Println(f("this is way.jpg"))
}

func day15Ex1() {
	str := "hello周四"
	runs := []rune(str)

	for i := 0; i < len(runs); i++ {
		fmt.Printf("%c \n", runs[i])
	}
}
