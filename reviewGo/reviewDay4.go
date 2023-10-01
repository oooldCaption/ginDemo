package reviewGo

import "fmt"

func RunDay4() {
	structCompare()
}

type person struct {
	name string
	age  int
}

func (p person) run() {
	fmt.Println("跑得很快", p.name)
}

func structCompare() {
	s1 := person{
		name: "oldSu",
		age:  30,
	}
	s2 := person{
		name: "oldSu",
		age:  30,
	}

	/// 结构体可以比较, 顺序.值都相等, 认为他们相等
	if s1 == s2 {
		fmt.Println("s1 == s2")
	}

	s3 := struct {
		name string
		age  int
	}{
		age:  33,
		name: "Q",
	}
	s4 := struct {
		age  int
		name string
	}{
		age:  33,
		name: "Q",
	}
	// s3 和 s4 不能比较, 不是同一种类型
	fmt.Println(s3)
	fmt.Println(s4)

}
