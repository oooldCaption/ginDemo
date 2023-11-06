package reviewGo

import (
	"ginDemo/reviewGo/model"
)

func RunDay19() {
	day19Ex2()
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
