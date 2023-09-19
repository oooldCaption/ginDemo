package reviewGo

import "fmt"

func RunReviewDayOne() {
	dayOne()
}

func dayOne() {
	slice := []int{1, 2, 3, 4, 5, 6}
	m := make(map[int]*int)

	for i, i2 := range slice {
		// i2 是一个临时变量, 在迭代的时候并不会改变
		// 所有的 键 都指向了 slice 最后一个值,6
		m[i] = &i2
	}

	// 观察打印信息, 可以看到 i2 的地址都是一样的.
	for i, i2 := range m {
		//fmt.Println(i, i2)
		fmt.Println(i, *i2)
	}

	// 修改
	for i, i2 := range slice {
		v := i2
		m[i] = &v
	}
	for i, i2 := range m {
		//fmt.Println(i, i2)
		fmt.Println(i, *i2)
	}
}
