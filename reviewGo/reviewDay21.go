package reviewGo

import (
	"fmt"
	"math"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

func RunDay21() {
	day21Ex6()
}

func day21Ex6() {
	ch := make(chan int)
	go producer(ch)
	go consumer(ch)
	time.Sleep(time.Second * 2)
}
func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}
func consumer(ch chan int) {
	for v := range ch {
		fmt.Println("Consumed: ", v)
	}
}

func day21Ex5() {
	ch := make(chan int)
	go findPrime(0, 9000000, ch)
	fmt.Println("程序开始时间:", time.Now())
	for _ = range ch {
		//fmt.Println(p)
	}
	fmt.Println("程序结束时间:", time.Now())
}
func findPrime(begin, end int, ch chan int) {
	for i := begin; i < end; i++ {
		if isPrime(i) {
			ch <- i
		}
	}
	close(ch)
}

func day21Ex4() {
	fmt.Println("程序开始时间", time.Now())
	for i := 0; i < 9000000; i++ {
		if isPrime(i) {
			//fmt.Println("素数:", i, time.Now())
		}
	}
	fmt.Println("程序结束时间", time.Now())
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func day21Ex1() {
	go day21Ex2()
	for i := 0; i < 10; i++ {
		fmt.Println("this is way")
		time.Sleep(time.Second)
	}
}
func day21Ex2() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello world")
		time.Sleep(time.Duration(time.Second))
	}
}

func day21Ex3() {
	var wg sync.WaitGroup
	resultChan := make(chan string)

	// Example range of IPs to ping, adjust your range accordingly
	for i := 0; i <= 255; i++ {
		for j := 0; j <= 255; j++ {
			ip := fmt.Sprintf("10.0.%d.%d", i, j)
			wg.Add(1)
			go pingIP(ip, &wg, resultChan)
		}
	}

	// Close the result channel when all pings are done
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Read from result channel as results come in
	for result := range resultChan {
		fmt.Println(result)
	}
}

func pingIP(ip string, wg *sync.WaitGroup, resultChan chan<- string) {
	defer wg.Done() // Ensure that the wait group counter is decremented when this function finishes

	var cmd *exec.Cmd

	// Depending on the OS, the ping command may differ, adjust as needed
	if runtime.GOOS == "windows" {
		cmd = exec.Command("ping", "-n", "1", "-w", "1000", ip)
	} else {
		cmd = exec.Command("ping", "-c", "1", "-W", "1", ip)
	}

	if err := cmd.Run(); err != nil {
		resultChan <- fmt.Sprintf("IP %s is not reachable: %s", ip, err)
		return
	}

	resultChan <- fmt.Sprintf("IP %s is reachable", ip)
}
