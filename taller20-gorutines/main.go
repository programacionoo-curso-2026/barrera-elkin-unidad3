package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go ShowGoroutines(i)
	}
	go ShowGoroutines(1)
	time.Sleep(1 * time.Second)
}
func ShowGoroutines(id int) {
	delay := rand.Intn(500)
	fmt.Printf("Goroutines #%d with %dms\n", id, delay)
	time.Sleep(time.Millisecond * time.Duration(delay))
}
