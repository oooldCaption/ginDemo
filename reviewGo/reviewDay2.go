package reviewGo

import "fmt"

func RunDay2() {
	newAndMake()
}

func newAndMake() {
	// new 和 make 都是内置函数, 用来分配内存. 但是适用的数据类型不同;
	// new 返回一个指针, 指针指向了 新分配的 类型为 T 的零值; 适用于 值类型, 数组, 结构体
	// make 返回一个初始化后引用. 适用于 slice map channel
	// 使用 new()
	p := new(int) // 分配 int 的零值，返回 *int 类型的指针
	fmt.Printf("p: %v, *p: %v, type of p: %T\n", p, *p, p)

	// 使用 make()
	s := make([]int, 3) // 创建一个有 3 个 int 零值元素的切片
	fmt.Printf("s: %v, type of s: %T\n", s, s)

	m := make(map[string]int) // 创建一个空的 map
	fmt.Printf("m: %v, type of m: %T\n", m, m)

	c := make(chan int) // 创建一个无缓冲的 int 类型 channel
	fmt.Printf("c: %v, type of c: %T\n", c, c)

	// 注意：下面的代码会出错，因为 new() 不支持切片、映射或通道。
	// slice := new([]int)
	// map1 := new(map[string]int)
	// ch := new(chan int)
}

/// 函数 同时具有命名和未命名的返回形参 '(sum int, error)'
func funcMui(x, y int) (sum int, err error) {
	return x + y, nil
}

func sliceLength() {

	// 查看初始化的时候是否设置了 length

	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	// [0 0 0 0 0 1 2 3]

	t := make([]int, 0)
	t = append(t, 3, 2, 1)
	fmt.Println(t)
	//[3 2 1]
}
