package testTime

import (
	"fmt"
	"time"
)

// time 包是 golang 中内置的关于 时间 日期的功能模块

func TimeLib() {
	//getCurrentTime()
	//getTimestamp()
	//timeDuration()
	testTick()
}

// 获取当前时间
func getCurrentTime() {
	now := time.Now()

	fmt.Println(now)
	fmt.Printf("%#v", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	min := now.Minute()
	sec := now.Second()
	fmt.Printf("\n%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, min, sec)

	// 要注意时间格式, 防止 mm 跟 MM 错位
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}

// 测试时间戳跟标准时间进行互相转换
func getTimestamp() {
	now := time.Now()
	unixTimestamp := now.Unix()
	unixTimestampNano := now.UnixNano()
	fmt.Println(unixTimestamp, unixTimestampNano)

	// timestamp to time
	t := time.Unix(unixTimestamp, 0)
	fmt.Println(t)
}

// 时间间隔 增 减 计算等
func timeDuration() {
	now := time.Now()

	// 时间累加
	latterFiveMin := now.Add(time.Minute * 5)
	fmt.Println(latterFiveMin)

	// 在 某个时间基础上累加 单位时间量度
	t2 := now.AddDate(1, 1, 1)
	fmt.Println(t2)

	// 计算 两个时间点的 duration
	t3 := t2.Sub(now)
	fmt.Println(t3)

	// 比较 两个时间是否相等. 会考虑时区
	isEq := t2.Equal(now)
	fmt.Println(isEq)

	// t2 的时间 是否在 now 之后
	isAfter := t2.After(now)
	fmt.Println(isAfter)

	isBefore := t2.Before(now)
	fmt.Println(isBefore)

}

// 测试定时器
func testTick() {
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Println(i)
	}
}
