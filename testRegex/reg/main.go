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
	testEndPix("demo.pdf")
}

func testEndPix(fileName string) {
	s0, r0 := "/OK/g", "Understand? OK or Not"
	s1, r1 := "demo.pdf", "^\\w+\\.pdf$"
	s2, r2 := "abcedf", "/[e-o]/g"

	tempRegList := []string{r0, r1, r2}
	tempStrList := []string{s0, s1, s2}

	for i := 0; i < len(tempRegList); i++ {
		fileEnd, err := regexp.MatchString(tempRegList[i], tempStrList[i])
		fmt.Println(fileEnd, err)
	}

}
