package reviewGo

import (
	"fmt"
	"sort"
)

func RunDay7() {
	//day7Ex1()
	//fmt.Println(test("qwertyuio", "qwe"))
	//fmt.Println(test("a", "a"))
	//fmt.Println(twoSum([]int{1, 2, 3, 4, 5}, 6))
	//fmt.Println(twoSum([]int{1, 2, 3, 4, 5}, 9))
	fmt.Println(lengthOfLongestSubstring("asdfghjklkkl"))
	fmt.Println(lengthOfLongestSubstring("abcdabc"))
}

// 滑动窗口解决:
// 滑动窗口是一个 解决数组/字符串 问题常见的方案;
// 使用 两个指针 left, right; 作为当前 子数组/子字符串的开始和结束;
// 随着问题的 解 进行滑动.
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	charMap := make(map[byte]int)
	maxLength := 0
	left, right := 0, 0

	for right = 0; right < n; right++ {
		// 如果 字符串出现重复, 移动 left 指针
		if _, ok := charMap[s[right]]; ok {
			left = max(charMap[s[right]]+1, left)
		}
		// 更新最长子串长度
		maxLength = max(maxLength, right-left+1)
		charMap[s[right]] = right
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 解题思路:
// 1. 循环目标数组
// 2. 将遍历结果存储到 map 中, 然后再 map 中寻找给定值的另一个数字
// 3. 存在的话返回这两个数字的下标
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		fmt.Println(m)
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	fmt.Println(m)
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumber(l1, l2 *ListNode) *ListNode {

	head := &ListNode{Val: 0}
	n1, n2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}

func test(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:len(needle)] == needle {
			return i
		}
	}
	return -1
}

// 查找子串的索引位置
func strStr(haystack string, needle string) int {

	// 当 子串为 0 的时候, 返回索引 0
	if len(needle) == 0 {
		return 0
	}

	// 定义游标
	var i, j int

	// 开始匹配是否有 与 needle 匹配的子串.
	for i = 0; i < len(haystack)-len(needle)+1; i++ {

		for j = 0; j < len(needle); j++ {
			// 如果出现不匹配. 跳出循环
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if len(needle) == j {
			return i
		}
	}

	return -1
}

type Rope struct {
	Duration    int
	BurningEnds int
}

//有两根不均匀的绳子, 每根绳子燃烧时间都为 1 小时. 用两根绳子的燃烧时间来设定一个 45min 的计时
func day7Ex1() {
	//将绳子A的两端同时点燃，同时将绳子B的一端点燃。由于绳子A从两端燃烧，所以它会在30分钟内燃烧完。
	//当绳子A燃烧完毕时，立即将绳子B的另一端也点燃。由于绳子B此时从两端燃烧，所以它会在接下来的15分钟内燃烧完，此时总共过去了45分钟。
	ropeA := Rope{
		Duration:    60,
		BurningEnds: 2,
	}
	ropeB := Rope{
		Duration:    60,
		BurningEnds: 1,
	}
	totalTime := 0
	for {
		totalTime++

		ropeA.Duration -= ropeA.BurningEnds
		ropeB.Duration -= ropeB.BurningEnds

		if ropeA.Duration <= 0 {
			ropeB.BurningEnds = 2
		}
		if ropeB.Duration <= 0 {
			fmt.Println(totalTime)
			break
		}
	}
}

//编写一个程序，找出一个整数切片中的最大值和最小值
func getMaxAndMin() {
	// 这样写也会被面试官打死
	sli := []int{11, 22, 323, 1214, 35}
	sort.Ints(sli)
	fmt.Println(sli[0])
	fmt.Println(sli[len(sli)-1])
	min, max := findMaxAndMin(sli)
	fmt.Println("Minimum:", min)
	fmt.Println("Maximum:", max)

}

func findMaxAndMin(sli []int) (min int, max int) {
	// 设置初始值
	min = sli[0]
	max = sli[0]

	for _, v := range sli {
		if v < min {
			min = v
		}
		for v > max {
			max = v
		}
	}
	return min, max
}
