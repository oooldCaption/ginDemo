package reviewGo

import "fmt"

func RunDay5() {
	day5Ex3()
}

type User struct {
	Name  string
	Email string
}

func (u *User) notify() {
	fmt.Println(u.Name, u.Email)
}

// 方法接受值类型和指针类型的差别
func (u *User) modify(name, email string) {
	u.Email = email
	u.Name = name
	fmt.Println(u.Name, u.Email)
}

func day5Ex1() {
	u := User{
		Name:  "oldSu",
		Email: "dipsy@xxx.com",
	}
	u.notify()
	u.modify("Z", "z.xxx")
	u.notify()
}

type Manager struct {
	User
}

func (self *User) ToString() string {
	return fmt.Sprintf("User: %p, %v", self, self)
}
func day5Ex2() {
	m := Manager{User{"ooo", "Tom"}}
	fmt.Printf("Manager: %p\n", &m)
	fmt.Println(m.ToString())
}

type Person interface {
	Speak(string) string
}
type Student struct {
}

func (s Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是大帅比"
	} else {
		talk = "您好"
	}
	return talk
}
func day5Ex3() {
	var p Person = Student{}
	think := "1"
	fmt.Println(p.Speak(think))
}
