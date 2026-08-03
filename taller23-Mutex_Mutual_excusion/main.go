package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
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

func main() {
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

func updateOrderStatus(order *Order) {
	order.mu.Lock()
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	status := []string{
		"Procesando", "Despachando", "Entregado",
	}[rand.Intn(3)]
	order.Status = status
	fmt.Printf("Actualizando orden %d con estado: %s\n",
		order.ID, status)
	order.mu.Unlock()

	updateMutex.Lock()
	defer updateMutex.Unlock()

	currentUpdates := totalUpdates
	time.Sleep(5 * time.Millisecond)
	totalUpdates = currentUpdates + 1

}

func generateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{
			ID: i + 1, Status: "pending",
		}
	}
	return orders
}
