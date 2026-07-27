package main

import (
	"fmt"
)

type Order struct {
	ID     int
	Status string
}

func main() {
	orders := generateOrders(20)
	fmt.Printf("Numeros de ordenes %d\n", len(orders))
	fmt.Print("Todas las operaciones completadas. Finalizando\n")
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
