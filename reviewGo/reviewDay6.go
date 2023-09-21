package reviewGo

import (
	"fmt"
	"math"
	"strings"
	"unsafe"
)

func RunDay6() {
	day6Ex5()
}

// 字符串反转
func day6Ex4() {
	s := "asdfghjkl"
	l := len(s) - 1
	var c string
	for i := l; i >= 0; i-- {
		c = c + string(s[i])
	}
	fmt.Println(c)
}

// 字符串反转 2
func day6Ex5() {
	a := reverseString("asdfghjkl")
	fmt.Println(a)
}

// 算是进阶解法吧
func reverseString(s string) string {
	/// rune 类型的切片
	runes := []rune(s)
	// 定义两个游标
	i, j := 0, len(runes)-1
	for i < j {
		// 前后交换
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)
}

// 检查一个字符串是否是另一个字符串的子串
func day6Ex3() {
	s := "asdfghjkl"
	t := "asd"

	/// 面试官可能会打死我
	fmt.Println(strings.Contains(s, t))
	fmt.Println(containsSubstring(s, t))
	fmt.Println(containsSubstring("asdfghjkl", "kl"))

}

func containsSubstring(s, t string) bool {
	//遍历字符串 s 直到 len(s) - len(t)。因为如果 s 的剩余部分比 t 短，t 就不可能是 s 的子串。
	//对于 s 的每一个字符，取出一个与 t 同样长度的子串，并与 t 进行比较。
	//如果在某个位置找到了一个与 t 相同的子串，那么返回 true。
	//如果遍历完整个 s 都没有找到与 t 相同的子串，那么返回 false。
	//这种方法的时间复杂度是 O(n*m)，其中 n 是字符串 s 的长度，m 是字符串 t 的长度。
	for i := 0; i <= len(s)-len(t); i++ {
		if s[i:i+len(t)] == t {
			return true
		}
	}
	return false
}

/// 交换两个变量
func swap() {
	x, y := 7, 8
	x, y = y, x
	fmt.Println(x, y)
}

/// 计算圆的周长和面积
func circleAreaAndPerimeter(r float64) {
	area := math.Pi * r * r
	circumference := 2 * math.Pi * r
	fmt.Printf("圆的半径:%f, 周长:%f, 面积:%f \n\n", r, circumference, area)
}

/// 空结构体模拟 Set 结构
func day6Ex1() {
	set := make(map[string]struct{})
	fmt.Println(unsafe.Sizeof(set))

	/// 空结构体不占内存, s 和 set 的 size 一致
	s := make(map[string]string)
	fmt.Println(unsafe.Sizeof(s))

	/// 空结构体模拟 set
	set["item1"] = struct{}{}
	set["item2"] = struct{}{}
	_, ok := set["item2"]
	if ok {
		fmt.Println("item2 exit")
	}
	fmt.Println(unsafe.Sizeof(set))
}
