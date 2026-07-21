package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go ShowGoroutines(1)
	time.Sleep(1 * time.Second)
}

func ShowGoroutines(id int) {
	delay := rand.Intn(500)
	fmt.Printf("Goroutines #%d\n with %ms", id, delay)
	time.Sleep(time.Millisecond * time.Duration(delay))
}
