package main

import (
	"fmt"
)

type Order struct {
	id       string
	items    []OrderItem
	status   string
	customer Customer
}

type OrderItem struct {
	name    string
	price   float64
	quanity int
}

type Customer struct {
	id   string
	name string
}

func (o *Order) GetTotal() float64 {
	var total float64
	for _, item := range o.items {
		total += item.price * float64(item.quanity)
	}
	return total
}
func (o *Order) AddItem(name string, price float64, quanity int) {
	item := OrderItem{
		name:    name,
		price:   price,
		quanity: quanity,
	}
	o.items = append(o.items, item)
}
func (o *Order) RemoveItem(name string) {
	for i, item := range o.items {
		if item.name == name {
			o.items = append(o.items[:i], o.items[i+1:]...)
			break
		}
	}
}
func (o *Order) SetStatus(status string) {
	o.status = status
}
func main() {
	customer := Customer{
		id:   "1",
		name: "Феодосий",
	}

	order := Order{
		id:       "00001",
		status:   "Новый",
		customer: customer,
	}

	order.AddItem("Парфюм(BlackAfgano)", 3200, 1)
	order.AddItem("Чай", 250, 2)
	order.AddItem("Майонез", 150, 1)

	fmt.Printf("Заказ #%s\n", order.id)
	fmt.Printf("Клиент: %s\n", order.customer.name)
	fmt.Printf("Статус: %s\n", order.status)
	fmt.Printf("Товары в заказе:\n")
	for _, item := range order.items {
		fmt.Printf("	%s: %v руб. x %d\n", item.name, item.price, item.quanity)
	}
	fmt.Printf("Итого: %.2f руб.\n", order.GetTotal())

	order.SetStatus("В обработке")
	fmt.Printf("Новый статус: %s\n", order.status)

	order.RemoveItem("Майонез")

	fmt.Printf("\nПосле удаления Майонеза:\n")

	fmt.Printf("Итого: %.2f руб.\n", order.GetTotal())
}
