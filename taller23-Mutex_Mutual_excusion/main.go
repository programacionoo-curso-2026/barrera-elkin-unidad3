package main

import (
	"fmt"
	"sync"
)

type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}

var (
	totalUpdates int
	updateMutex  sync.Mutex
)

func main_n() {
	var wg sync.WaitGroup
	wg.Add(3)
	orders := generateOrders(20)
	//go func() {
	//  defer wg.Done()
	//  processOrders(orders)
	//}()
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			for _, order := range orders {
				updateOrderStatus(order)
			}
		}()
	}
	//reportOrderStatus(orders)

	wg.Wait()

	fmt.Print("Todas las operaciones completadas. Saliendo\n")
	fmt.Printf("Total Actualizaciones %d\n", totalUpdates)
}
