package model

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
