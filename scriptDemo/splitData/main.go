package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
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
	sourceDir string = "./"
)

func main() {
	fmtFiles()
}

func fmtFiles() {
	fmt.Println(time.Now())
	createRootDir()
	getFileList()
	fmt.Println("运行结束")
	fmt.Println(time.Now())
}

//0.041813
//0.056368
func getFileList() {
	files, _ := ioutil.ReadDir(sourceDir)
	for _, file := range files {
		fp := file
		if strings.HasSuffix(fp.Name(), "txt") {
			go fmtFile(fp)

		}
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

	os.Create(outputFileName)
	writeFile, wErr := os.OpenFile(outputFileName, os.O_RDWR|os.O_APPEND, 0666)
	if wErr != nil {
		panic(wErr)
	}
	defer writeFile.Close()
	write := bufio.NewWriter(writeFile)

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if strings.Contains(line, "(0x) 03-00-00-02-08-00-00-04-02") {
			hexStr := line[len(line)-77 : len(line)]

			x := returnXYZ(hexToInt(strings.Split(hexStr, "-")[10] + strings.Split(hexStr, "-")[9]))
			x1 := returnXYZ(hexToInt(strings.Split(hexStr, "-")[16] + strings.Split(hexStr, "-")[15]))

			y := returnXYZ(hexToInt(strings.Split(hexStr, "-")[12] + strings.Split(hexStr, "-")[11]))

			y1 := returnXYZ(hexToInt(strings.Split(hexStr, "-")[18] + strings.Split(hexStr, "-")[17]))

			z := returnXYZ(hexToInt(strings.Split(hexStr, "-")[14] + strings.Split(hexStr, "-")[13]))
			z1 := returnXYZ(hexToInt(strings.Split(hexStr, "-")[20] + strings.Split(hexStr, "-")[19]))

			write.WriteString("x:" + strconv.FormatInt(x, 10) + ",y:" + strconv.FormatInt(y, 10) + ",z:" + strconv.FormatInt(z, 10) + "\n")
			write.WriteString("x:" + strconv.FormatInt(x1, 10) + ",y:" + strconv.FormatInt(y1, 10) + ",z:" + strconv.FormatInt(z1, 10) + "\n")

		}
		write.Flush()

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

func returnXYZ(value int64) int64 {
	if value <= 32767 {
		return value
	} else {
		var result int64
		str := strconv.FormatInt(value, 2)
		tempStr := ""
		for i := 0; i < len(str); i++ {
			if str[i:i+1] == "1" {
				tempStr += "0"
			} else {
				tempStr += "1"
			}
		}
		result, _ = strconv.ParseInt(tempStr, 2, 32)
		return -result
	}
}
func hexToInt(value string) int64 {
	result, _ := strconv.ParseInt(value, 16, 32)
	return result
}

// 创建顶部目录
func createRootDir() {
	currentDate := time.Now()
	if isTest {
		rootDir, _ = os.Getwd()
	} else {
		rootDir, _ = os.Executable()
		rootDir = rootDir[0 : len(rootDir)-5]
	}
	sourceDir = rootDir + "/" + fmt.Sprintf("%02d", currentDate.Month()) + "-" + fmt.Sprintf("%02d", currentDate.Day())
	rootDir += "/output-" + fmt.Sprintf("%02d", currentDate.Month()) + "-" + fmt.Sprintf("%02d", currentDate.Day())
	fmt.Println("从中读取", sourceDir, ", 写入到", rootDir)
	os.Mkdir(rootDir, os.ModePerm)
}

func textHex() {

	hexStr := "03-00-00-02-08-00-00-04-02-00-00-07-00-05-FF-00-00-06-00-07-FF-00-00-00-00-00"
	x := returnXYZ(hexToInt(strings.Split(hexStr, "-")[10] + strings.Split(hexStr, "-")[9]))
	x1 := returnXYZ(hexToInt(strings.Split(hexStr, "-")[16] + strings.Split(hexStr, "-")[15]))

	y := returnXYZ(hexToInt(strings.Split(hexStr, "-")[12] + strings.Split(hexStr, "-")[11]))

	y1 := returnXYZ(hexToInt(strings.Split(hexStr, "-")[18] + strings.Split(hexStr, "-")[17]))

	z := returnXYZ(hexToInt(strings.Split(hexStr, "-")[14] + strings.Split(hexStr, "-")[13]))
	z1 := returnXYZ(hexToInt(strings.Split(hexStr, "-")[20] + strings.Split(hexStr, "-")[19]))

	fmt.Println(x, y, z)
	fmt.Println(x1, y1, z1)
}
