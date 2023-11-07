package reviewGo

import (
	"fmt"
	"ginDemo/reviewGo/model"
	"os/exec"
	"strconv"
	"strings"
)

func RunDay19() {
	day19Ex3()
}

func day19Ex1() {
	phone := model.Phone{}
	var usbPhone model.USB = phone
	phone.Start()
	usbPhone.Start()

	if p, ok := usbPhone.(model.Phone); ok {
		p.PhoneCall()
	}
}

func day19Ex5() {
	/// map 遍历是无序的
	tempMap := make(map[int]int)
	tempMap[1] = 1
	tempMap[2] = 2
	tempMap[3] = 3
	tempMap[4] = 4
	tempMap[5] = 5
	for i, i2 := range tempMap {
		fmt.Println(i, i2)
	}
}

func day19Ex2() {
	var lielie model.CustomHero = model.SuLie{
		Name:  "苏烈",
		Level: 15,
	}
	var popo model.CustomHero = model.LianPo{
		Name:  "廉颇",
		Level: 15,
	}
	var xiangyu model.CustomHero = model.XiangYu{
		Name:  "项羽",
		Level: 15,
	}
	var xiaojinjin model.CustomHero = model.ChengYaoJin{
		Name:  "程咬金",
		Level: 15,
	}
	var heroList = []model.CustomHero{lielie, popo, xiangyu, xiaojinjin}
	p := model.Player{}
	for _, hero := range heroList {
		p.Play(hero)
	}
}

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

var padingStr = "------------"
var spaceStr = "             "

// day19Ex3  简易版家庭记账本
func day19Ex3() {
	println(ColorCyan, padingStr+"家庭记账系统"+padingStr, ColorReset)
	println(spaceStr + "1.查看收支明细")
	println(spaceStr + "2. 存入")
	println(spaceStr + "3. 支出")
	println(spaceStr + "4. 退出")

	inputCMD()
}

func inputCMD() {
	var cmd string
	_, err := fmt.Scanln(&cmd)
	if err != nil {
		panic("输入异常")
		return
	}
	switch cmd {
	case "1":
		showAccountDetail()
	case "2":
		addMoney()
	case "3":
		costMoney()
	case "4":
		exitAccount()
	default:
		println("输入不符合预期, 请重试\n")
		day19Ex3()
	}
}

func showAccountDetail() {

}
func addMoney() {

}

func costMoney() {

}
func exitAccount() {

}

// test terminal support colorful
func day19Ex4() {
	cmd := exec.Command("tput", "colors")
	output, err := cmd.Output()
	if err == nil {
		colors, _ := strconv.Atoi(strings.TrimSpace(string(output)))
		if colors >= 256 {
			fmt.Println("Your terminal supports 256 colors!")
		}
	}
}
