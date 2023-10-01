package reviewGo

import (
	"fmt"
	"time"
)

func RunDay9() {
	day9Ex5()
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		maxLen := max(len1, len2)
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	L, R := left, right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}

// 函数测试

type MathFunc func(a, b int) int

func applyMathFunc(mf MathFunc, a, b int) int {
	return mf(a, b)
}
func add(a, b int) int {
	return a + b
}

func day9Ex2() {
	var mf MathFunc
	mf = add
	result := applyMathFunc(mf, 1, 2)
	fmt.Println(result)
}

// OCP 原则
type Pet interface {
	eat()
	//sleep()
	//run()
	//play()
}

type Dog struct {
	name string
}

func (d Dog) eat() {
	fmt.Println("dog eat", d.name)
}

type Cat struct {
	name string
}

func (c Cat) eat() {
	fmt.Println("cat eat", c.name)
}

type Bird struct {
	name string
}

func (b Bird) eat() {
	fmt.Println("bird eat", b.name)
}

type Ow struct {
	name string
}

func (o Ow) cate(pet Pet) {
	pet.eat()
}

func day9Ex3() {
	d := Dog{name: "狗子"}
	c := Cat{name: "猫子"}
	p := Ow{name: "laowang"}
	p.cate(d)
	p.cate(c)
}

func day9Ex4() {
	p := person{
		name: "33",
		age:  0,
	}
	p.run()
}

func day9Ex5() {
	go show("test")
	//show("j")
}

func show(str string) {
	for i := 0; i < 5; i++ {
		fmt.Println("msg: %v \n", str)
		time.Sleep(time.Millisecond * 100)
	}
}

func testGIF() {
	//gif.Decode()
}
