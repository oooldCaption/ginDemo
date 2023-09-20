package reviewGo

func RunDay3() {
	ex3()
}
func ex1() {
	//list := new([]int)
	// 在昨天的一个例子中, 我们提到 new 返回的结果是一个 *[]int 的指针, 不能进行 append 操作;
	//list = append(list, 1)

	// 可以使用 make 初始化
	//l := make([]int, 5)
	//l = append(l, 1, 2, 3)
}

func ex2() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}

	//s1 = append(s1,s2) 错误的写法
	s1 = append(s1, s2...) // 使用展开操作符追加
}

// TIP: 简短声明模式:
//1.必须使用显示初始化；
//2.不能提供数据类型，编译器会自动推导；
//3.只能在函数内部使用简短模式；

//var (
//	size := 23
//	max_size = size*2
//)
func ex3() {
	var (
		size     = 23
		max_size = size * 2
	)
	print(size, max_size)
}
