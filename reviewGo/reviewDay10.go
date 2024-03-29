package reviewGo

import (
	"cmp"
	"fmt"
	"sync"
)

func RunDay10() {
	//fmt.Println(maxValue[int](1, 3))
	//fmt.Println(maxValue(1.2, 1.4))
	//testData()
	day10Ex3()
}
func day10Ex3() {
	v := make(chan string)
	<-v
	fmt.Println("this end")

}

func day10Ex2() {
	var sw sync.WaitGroup
	r := make(chan struct{})

	for i := 0; i < 10; i++ {
		sw.Add(1)
		go func(id int) {
			defer sw.Done()
			<-r
			println(id)
		}(i)
	}
	close(r)
	sw.Wait()
}
func day10Ex1() {
	go func() {
		defer fmt.Println("协程进行结束")
	}()
	defer fmt.Println("main end")
}

func maxValue[T cmp.Ordered](x, y T) T {
	if x > y {
		return x
	}
	return y
}

type Data[T any] struct {
	x T
}

func (d Data[T]) test() {
	fmt.Println(d)
}

func testData() {
	x := Data[int]{x: 10}
	x.test()

}

func testWait() {
	var waitGroup sync.WaitGroup
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go hello(i, &waitGroup)
	}

	waitGroup.Wait()

}
func hello(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello done", i)
}
