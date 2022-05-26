package testFmt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
* 	fmt 库实现了类似 C 的printf scanf 等功能的 I/O. 主要包含 输出  & 输入 两个部分
 */

func TestPrint() {
	fmt.Print("直接输入内容 \n ")
	name := "煎人寿"
	age := 29
	fmt.Printf("my name is %s , i'am %d years old \n", name, age)
	fmt.Println("自带换")
	fmt.Println("行")
}

// Fprint 系列的函数会将内容除数刀 io/write接口类型的变量中. 可以写入文件
func TestFPrint() {
	//fmt.Fprint(os.Stdout, "文件写入内容")
	fileObj, err := os.OpenFile("./fprint.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println("文件操作失败.", err)
		return
	}
	name := "煎人寿"
	fmt.Fprint(fileObj, "log message: ", name)
}

// Sprint 系列的函数会将传入的数据生成 并返回为一个字符串
func TestSprint() {
	s1 := fmt.Sprintln("煎人寿")
	name := "ooooldCaption"
	age := 29
	s2 := fmt.Sprintf("name is :%s , age is : %d", name, age)
	s3 := fmt.Sprintln("this is test msg")
	fmt.Println(s1, s2, s3)
	fmt.Println("------")
}

// scan 从标准输入中扫描文本
func TestScan() {
	testBufio()
	return
	var (
		name string
		age  int
	)
	fmt.Println("请输入姓名 年龄:")
	//fmt.Scanln(&name, &age)

	// 1:煎人寿 2:29  -> 接受信息: 煎人寿 29
	//fmt.Scanf("1:%s 2:%d", &name, &age)

	/// 根据换行界定输入是否结束
	//fmt.Scanln(&name, &age)

	fmt.Println("接受信息: ", name, age)
}

// 获取完整内容. 根据一个特定的 标记 结束读取
func testBufio() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入:")
	// p 为结束标记
	text, _ := reader.ReadString('p')
	text = strings.TrimSpace(text)
	fmt.Println()
	fmt.Println(text)
	fmt.Println()

	fmt.Printf("%#v", text)
}
