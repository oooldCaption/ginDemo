package reviewGo

import "fmt"

func RunDay14() {
	//println(day14Ex3(15))
	//println(day14Ex4(1))
	//println(day14Ex4(4))
	println(day14Ex5(1))
}

// 猴子吃桃子
// 有一堆桃子, 第一天吃其中的一半,并且再多吃一个, 以后的每一天, 猴子吃其中的一半, 并且多一个; 当第十天的时候(还没吃) 发现只剩一个桃子;
// 一共多少桃子
func day14Ex5(n int) int {
	if n > 10 || n < 1 {
		return 0
	}
	if n == 10 {
		return 1
	} else {
		return (day14Ex5(n+1) + 1) * 2
	}
}

// 在递归深入到下一层之前，当前层的逻辑会暂停，并等待下一层递归返回结果。
func day14Ex4(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*day14Ex4(n-1) + 1
	}
}

func day14Ex3(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return day14Ex3(n-1) + day14Ex3(n-2)
	}
}

func day14Ex2(n int) {
	if n > 2 {
		n--
		day14Ex2(n)
	}
	println(n)
}

func day14Ex1() {
	//var str string = "hello world"
	//for i := 0; i < len(str); i++ {
	//	fmt.Println(str[i])
	//}
	//for i, i2 := range str {
	//	fmt.Println(i, i2)
	//}

	total := 0
	for i := 0; i < 100; i++ {
		if i%9 == 0 {
			total += i
		}
	}

	//for {
	//	total++
	//
	//	if total%9 == 0 {
	//		continue
	//	}
	//	println(total)
	//	if total >= 1000 {
	//		break
	//	}

	//total++
	//println(total)
	//}
	n := 15
	for i := 1; i <= n; i++ {
		// 打印空格
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}

		// 打印星星
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
