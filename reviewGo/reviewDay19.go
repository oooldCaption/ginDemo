package reviewGo

import "ginDemo/reviewGo/model"

func RunDay19() {
	day19Ex1()
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

}
