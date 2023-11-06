package model

import "fmt"

type USB interface {
	Start()
	Read()
	Write()
	Exit()
}

type Phone struct{}

func (p Phone) PhoneCall() {
	fmt.Println("call  some one")
}

func (p Phone) Start() {
	fmt.Println("phone start usb")
}
func (p Phone) Read() {
	fmt.Println("phone read usb")
}
func (p Phone) Write() {
	fmt.Println("phone write usb")
}
func (p Phone) Exit() {
	fmt.Println("phone exit usb")
}

type Pad struct{}

func (p Pad) Start() {
	fmt.Println("Pad start usb")
}
func (p Pad) Read() {
	fmt.Println("Pad read usb")
}
func (p Pad) Write() {
	fmt.Println("Pad write usb")
}
func (p Pad) Exit() {
	fmt.Println("Pad exit usb")
}

type Camera struct {
}

func (p Camera) Start() {
	fmt.Println("Camera start usb")
}
func (p Camera) Read() {
	fmt.Println("Camera read usb")
}
func (p Camera) Write() {
	fmt.Println("Camera write usb")
}
func (p Camera) Exit() {
	fmt.Println("Camera exit usb")
}

type CustomHero interface {
	UpLevel()
	Attack()
	Status()
}

type SuLie struct {
	Name  string
	Level int
}

func (s SuLie) UpLevel() {
	println(s.Name + "升级了")
}

func (s SuLie) Attack() {
	println(s.Name + "放大了")
}

func (s SuLie) Status() {
	println(s.Name + "状态不好")
}

type LianPo struct {
	Name  string
	Level int
}

func (l LianPo) UpLevel() {
	println(l.Name + "升级了")
}

func (l LianPo) Attack() {
	println(l.Name + "放大了")
}

func (l LianPo) Status() {
	println(l.Name + "状态不好")
}

type XiangYu struct {
	Name  string
	Level int
}

func (x XiangYu) UpLevel() {
	println(x.Name + "升级了")
}

func (x XiangYu) Attack() {
	println(x.Name + "放大了")
}

func (x XiangYu) Status() {
	println(x.Name + "状态不好")
}

type ChengYaoJin struct {
	Name  string
	Level int
}

func (c ChengYaoJin) UpLevel() {
	println(c.Name + "升级了")
}

func (c ChengYaoJin) Attack() {
	println(c.Name + "放大了")
}

func (c ChengYaoJin) Status() {
	println(c.Name + "状态不好")
}

type Player struct{}

func (p Player) Play(hero CustomHero) {

	switch hero.(type) {
	case SuLie:
		hero.Attack()
	case LianPo:
		hero.Attack()
	case XiangYu:
		hero.Attack()
	case ChengYaoJin:
		hero.Attack()
	}
}
