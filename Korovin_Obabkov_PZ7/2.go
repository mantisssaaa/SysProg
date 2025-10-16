package main

import (
	"fmt"
)

func main() {
	warehouse := Inventory{}

	AddProduct(&warehouse, Product{ID: "1", name: "Хлеб", Price: 30, Quantity: 100})
	AddProduct(&warehouse, Product{ID: "2", name: "Колбаса", Price: 90, Quantity: 50})

	fmt.Printf("Общая стоимость: %.2f\n", GetTotalValue(warehouse))
	if test := WriteOff(&warehouse, "1", 50); test != nil {
		fmt.Println(test)
	}
	fmt.Printf("Общая стоимость после списания: %.2f\n", GetTotalValue(warehouse))
	if test := WriteOff(&warehouse, "2", 100); test != nil {
		fmt.Println(test)
	}
	if test := RemoveProduct(&warehouse, "3"); test != nil {
		fmt.Println(test)
	}
	fmt.Printf("%v\n", RemoveProduct(&warehouse, "99"))
}

type Product struct {
	ID, name string
	Price    float64
	Quantity int
}

type Inventory struct {
	products []Product
}

func AddProduct(inventory *Inventory, product Product) {
	inventory.products = append(inventory.products, product)
}

func WriteOff(inventory *Inventory, productID string, quantity int) error {
	for idx, p := range inventory.products {
		if p.ID == productID {
			if p.Quantity >= quantity {
				inventory.products[idx].Quantity -= quantity
				fmt.Printf("успешно списано %v %v\n", quantity, p.name)
				return nil
			}
			return fmt.Errorf("недостаточно товаров для такого списания")
		}
	}
	return fmt.Errorf("продукт с ID %s не найден", productID)
}

func RemoveProduct(inventory *Inventory, productID string) error {
	for i, p := range inventory.products {
		if p.ID == productID {
			inventory.products = append(inventory.products[:i], inventory.products[i+1:]...)
			fmt.Printf("успешно удален %v\n", p.name)
			return nil
		}
	}
	return fmt.Errorf("продукт с ID %s не найден", productID)
}

func GetTotalValue(inventory Inventory) float64 {
	var total float64
	for _, p := range inventory.products {
		total += p.Price * float64(p.Quantity)
	}
	return total
}
