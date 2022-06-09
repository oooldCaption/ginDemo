package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

/*
1. 根据文件名分割:  动作 - 用户 - 是否连续
eg: cd x l =>  侧蹲 - x - 连续

2. 逐行读取 nrf Connect 的log文件, 然后将符合条件的数据 处理为关键数据后 放进数组中
eg: I	13:28:13.836	Notification received from 6e400003-b5a3-f393-e0a9-e50e24dcca9e, value: (0x) 03-00-00-02-08-00-00-04-02-CF-00-78-FF-34-00-CC-00-78-FF-33-00-00-00-00-00-00
	根据解析规则解析 03-00-*****... 开头到结尾的十六进制数据

3. 将 2 中的数据 从hex 转换成 10 进制数据

4. 根据原始文件名生成 xyz三轴的log文件

5. 程序根据当前时间 创建对应的文件夹, 并将生成的文件放入到对应的文件夹中
eg: dd-hh-mm 文件夹名格式

*/

var (
	isTest    bool = true
	rootDir   string
	sourceDir string = "./source-06-08"
)

func main() {
	createRootDir()
	getFileList()

	return

}

func getFileList() {
	files, _ := ioutil.ReadDir(sourceDir)
	for _, file := range files {
		fmtFile(file)

	}
}
func fmtFile(info os.FileInfo) {
	// 创建文件
	filePrefix := strings.ReplaceAll(info.Name(), " ", "")
	filePrefix = filePrefix[0 : len(filePrefix)-4]
	var user string

	if filePrefix[2:3] == "s" {
		user = "su"
	} else if filePrefix[2:3] == "x" {
		user = "xia"
	} else if filePrefix[2:3] == "w" {
		user = "wang"
	}
	outputFileName := rootDir + "/" + user + "-" + strings.ToLower(filePrefix[0:2])
	if len(filePrefix) > 3 {
		outputFileName += "-lianxu.log"
	} else {
		outputFileName += ".log"
	}

	filePath := sourceDir + "/" + info.Name()

	file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if strings.Contains(line, "(0x) 03-00-00-02-08-00-00-04-02") {
			hexStr := line[len(line)-77 : len(line)]
		} else {
			return
		}

		if err != nil {
			if err == io.EOF {
				//fmt.Println(filePath, "文件读取完成")
				break
			} else {
				fmt.Println(filePath, "文件读取错误")
				return

			}
		}
	}

}

func returnXYZ() {

}
func hexToInt(value string) int {
	var result int = 0

	for _, v := range value {
		hexDigit := hex.
	}
	return result
}

// 创建顶部目录
func createRootDir() {
	currentDate := time.Now()
	if isTest {
		rootDir, _ = os.Getwd()
	} else {
		rootDir, _ = os.Executable()
	}
	rootDir += "/" + fmt.Sprintf("%02d", currentDate.Month()) + "-" + fmt.Sprintf("%02d", currentDate.Day())
	os.Mkdir(rootDir, os.ModePerm)
}
