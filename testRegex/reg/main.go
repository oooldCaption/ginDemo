package main

import (
	"fmt"
	"regexp"
)

//
//type regStruct struct {
//	rexStr string
//	rexCode string
//}

/// 代码中的正则表达式均来自 : https://regexlearn.com/zh-cn/learn/regex101
func main() {
	//testEndPix()
	createBinaryReg()
}

func createBinaryReg() {
	checkOutWithReg := func(regStr, str string) bool {
		reg, err := regexp.Compile(regStr)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return reg.MatchString(str)
	}
	print(checkOutWithReg("^\\w+\\.pdf$", "demo.pdf"))
}

func testEndPix() {
	s0, r0 := "/OK/g", "Understand? OK or Not"
	s1, r1 := "demo.pdf", "^\\w+\\.pdf$"
	s2, r2 := "abcdefghijklmn", "/[e-o]/g"

	tempRegList := []string{r0, r1, r2}
	tempStrList := []string{s0, s1, s2}

	for i := 0; i < len(tempRegList); i++ {
		fileEnd, err := regexp.MatchString(tempRegList[i], tempStrList[i])
		fmt.Println(fileEnd, err)
	}

}
