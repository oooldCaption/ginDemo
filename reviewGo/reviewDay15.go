package reviewGo

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func RunDay15() {
	day15Ex4()
}

func day15Ex4() {
	err := customErr("sd")
	if err != nil {
		panic(err)
	}
}
func testError() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error = ", err)
		}
	}()
	num1 := 10
	num2 := 0
	fmt.Println(num1 / num2)
}

func customErr(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("文件读取错误, 文件不存在")
	}
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

// 闭包 就是 一个函数 和 其相关的引用环境 组成的整体
/*
一种不严谨地解释:
可以把 闭包 当做一个类
函数是 操作
引用的环境 是字段
进行初始化后, 每次引用, 引用环境的状态会跟随变化
*/
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
	testFile()
}

func testFile() {
	f, _ := os.Open("reviewDay15.go")
	fmt.Println(f.Name())
	f.Seek(3, 0)
	buf := make([]byte, 10)
	n, _ := f.Read(buf)
	fmt.Println(n)
	fmt.Println(string(buf))
	defer f.Close()
}

func day15Ex1() {
	str := "hello周四"
	runs := []rune(str)

	for i := 0; i < len(runs); i++ {
		fmt.Printf("%c \n", runs[i])
	}
}
