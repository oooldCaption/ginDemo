package reviewGo

import (
	"fmt"
	"math"
)

func RunDay23() {
	checkOddEven(22)
	calSUmDivisibleByThree()
	ninenine()
}

func checkOddEven(num int) {
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}

func calSUmDivisibleByThree() {
	sum := 0
	for i := 1; i < 100; i++ {
		if i%3 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}

func ninenine() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i, "x", j, "=", i*j, "\t")
		}
		fmt.Println()
	}
}

func checkPrimeNumber(num int) {
	if num <= 1 {
		fmt.Println(num, "不是质数")
		return
	}
	if num == 2 {
		fmt.Println(num, "是质数")
		return
	}

	sqrtNum := int(math.Sqrt(float64(num)))
	for i := 2; i <= sqrtNum; i++ {
		if num%i == 0 {
			fmt.Println(num, "不是质数")
			return
		}
	}

	fmt.Println(num, "是质数")
}
