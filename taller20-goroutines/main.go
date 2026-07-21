package main

import (
	"fmt"
	"time"
)

func main() {
	go ShowGoroutines(1)
	time.Sleep(10 * time.Second)
}

func ShowGoroutines(id int) {
	fmt.Printf("Goroutines #%d\n", id)
}
