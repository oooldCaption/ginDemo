package model

import (
	"fmt"
)

type smallCat struct {
	Name  string
	age   int
	Width float32
}

// SetAge 可以通过结合 工厂模式 和 结构体的方法来控制访问权限
func (cat *smallCat) SetAge(age int) {
	cat.age = age
}

func NewSmallCat(name string, age int, width float32) *smallCat {
	return &smallCat{
		Name:  name,
		age:   age,
		Width: width,
	}
}

type Box struct {
	Long   float64
	Height float64
	Weight float64
}

func (box Box) BoxVolume() float64 {
	return box.Weight * box.Long * box.Height
}
func (box Box) BoxArea() float64 {
	return 2*box.Weight*box.Height + 2*box.Weight*box.Long + 2*box.Height*box.Long
}

type Visitor struct {
	Name  string
	Age   int
	Score int
}

func (v *Visitor) GetTicketPrice() int {
	if v.Age > 18 {
		v.Score = 30
	} else {
		v.Score = 0
	}
	return v.Score
}

type BankAccount struct {
	CardNum int
	Psw     string
	Balance float64
}

// SaveMoney 存钱业务
func (card *BankAccount) SaveMoney(money float64, psw string) {
	if psw != card.Psw {
		fmt.Println("密码错误")
		return
	}
	if money < 0 {
		fmt.Println("输入格式错误")
		return
	}
	fmt.Println("业务执行前卡片余额:", card.Balance)
	card.Balance += money
	fmt.Println("业务执行后卡片余额:", card.Balance)
}

// WithDraw 取钱业务
func (card *BankAccount) WithDraw(money float64, psw string) {
	if psw != card.Psw {
		fmt.Println("密码错误")
		return
	}
	if money < 0 {
		fmt.Println("输入格式错误")
		return
	}
	fmt.Println("业务执行前卡片余额:", card.Balance)
	card.Balance -= money
	fmt.Println("业务执行后卡片余额:", card.Balance)
}

// QueryCard 查询业务
func (card *BankAccount) QueryCard(psw string) {
	if psw != card.Psw {
		fmt.Println("密码错误")
		return
	}
	fmt.Println("卡片余额:", card.Balance)
}

type Student struct {
	Name  string
	Age   string
	ID    int
	Score float32
}

// ShowInfo 展示学生信息
func (stu *Student) ShowInfo() {
	fmt.Printf("学生信息: \n 姓名: %v \n 年龄: %v \n 学号: %v  \n 分数: %v\n", stu.Name, stu.Age, stu.ID, stu.Score)
}

// SetScore 设置分数
func (stu *Student) SetScore(score float32) {
	stu.Score = score
}

func (stu *Student) Say() string {
	return fmt.Sprintf("学生信息: \n 姓名: %v \n 年龄: %v \n 学号: %v  \n 分数: %v", stu.Name, stu.Age, stu.ID, stu.Score)
}

type Pupil struct {
	Student
}

func (stu *Pupil) testing() {
	fmt.Println("小学考试")
}

type Graduate struct {
	Student
}

func (stu *Graduate) testing() {
	fmt.Println("大学生考试")
}
