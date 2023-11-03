package reviewGo

import (
	"fmt"
	"ginDemo/reviewGo/model"
)

func RunDay18() {
	day18Ex5()
}

func day18Ex5() {
	card1 := model.BankAccount{
		CardNum: 111122223333,
		Psw:     "1234",
		Balance: 0,
	}
	card1.QueryCard("1234")
	card1.SaveMoney(-1, "1234")
	card1.SaveMoney(9991, "1232")
	card1.SaveMoney(9991, "1234")
	card1.QueryCard("1234")

	//card2 = model.BankAccount{
	//	CardNum: 444422221111,
	//	Psw:     "4321",
	//	Balance: 0,
	//}
	//card1.QueryCard("1222")

}

func day18Ex4() {
	v1 := model.Visitor{
		Name: "张大三",
		Age:  23,
	}
	v2 := model.Visitor{
		Name: "张老三",
		Age:  67,
	}
	v3 := model.Visitor{
		Name: "张小小三",
		Age:  11,
	}
	fmt.Println(v1.GetTicketPrice())
	fmt.Println(v2.GetTicketPrice())
	fmt.Println(v3.GetTicketPrice())
}

func day18Ex3() {
	box := model.Box{
		Long:   30.0,
		Height: 20,
		Weight: 22,
	}
	fmt.Println("体积: ", box.BoxVolume())
	fmt.Println("面积: ", box.BoxArea())
}

func day18Ex2() {
	stu := model.Student{
		Name:  "张三",
		Age:   "16",
		ID:    2020030101,
		Score: 98,
	}
	fmt.Println(stu.Say())
}
func day18Ex1() {
	cat := model.NewSmallCat("小花花", 23, 8.2)
	fmt.Println(cat)
	cat.SetAge(4)
	fmt.Println(cat)
}

func findCat() {
	cat1 := map[string]string{"name": "小白", "age": "5"}
	cat2 := map[string]string{"name": "小灰", "age": "3"}

	carArr := [2]map[string]string{cat1, cat2}
	name := "小白"
	for _, m := range carArr {
		if m["name"] == name {
			print(m["age"])
		}
	}
}
func findCatWithStruct() {
	type Cat struct {
		name  string
		age   int
		color string
	}
	cat1 := Cat{
		name:  "小花花",
		age:   20,
		color: "黑白色",
	}
	cat2 := Cat{
		name:  "小灰",
		age:   22,
		color: "灰色",
	}
	fmt.Println(cat1)
	fmt.Println(cat2)

}

func day18Ex0() {
	type Student struct {
		Name   string
		AgePtr *int
		Class  int
	}
	//s1 := Student{
	//	Name:   "sony",
	//	AgePtr: nil,
	//	Class:  8,
	//}
	//age := 28
	//s1.AgePtr = &age
	//fmt.Println(*s1.AgePtr)
	//age = 39
	//fmt.Println(*s1.AgePtr)

	s2 := &Student{
		Name:   "tt",
		AgePtr: nil,
		Class:  22,
	}
	fmt.Println(s2)

}
