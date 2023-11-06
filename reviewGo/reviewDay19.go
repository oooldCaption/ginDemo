package reviewGo

import (
	"fmt"
	"ginDemo/reviewGo/model"
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

var padingStr = "------------"

// day19Ex3  简易版家庭记账本
func day19Ex3() {
	println(padingStr + "家庭记账系统" + padingStr)
	println("请输入指令")
	var cmd string
	fmt.Scanf(cmd)
}
