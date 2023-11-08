package reviewGo

import (
	"fmt"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

func RunDay21() {
	day21Ex3()
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
