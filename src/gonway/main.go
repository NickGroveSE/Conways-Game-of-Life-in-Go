package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Second)
	var stepCount int = 0
	defer ticker.Stop()
	done := make(chan bool)

	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case <-ticker.C:
			stepCount++
			fmt.Println("Count: ", stepCount)
		}
	}

}
