package main

import "fmt"

type Product struct {
	name     string
	price    float64
	quantity int
}

func inventory() {

	inventory := make(map[string]Product)
	AddProduct(inventory, "Test-nuevo", 100.0, 50)
	AddProduct(inventory, "Producto", 100.0, 50)
	AddProduct(inventory, "Producto-1", 100.0, 50)

	UpdateQuantity(inventory, "Producto", 12)

	DeleteProduct(inventory, "Test-nuevo")

	ShowInventory(inventory)
}

func AddProduct(inventory map[string]Product, name string, price float64, quantity int) {

	newProduct := Product{name: name, price: price, quantity: quantity}

	inventory[name] = newProduct

	fmt.Println("New product was added")
}

func UpdateQuantity(inventory map[string]Product, name string, quantity int) {
	product, exist := inventory[name]
	if exist {
		product.quantity = quantity
		inventory[name] = product
		fmt.Printf("Product quantity '%s' updated to %d .\n ", name, quantity)
	} else {
		fmt.Printf("The Product '%s' doesnt exist. \n", name)
	}
}

func DeleteProduct(inventory map[string]Product, name string) {
	// Verificar si el producto existe en el inventario
	_, exist := inventory[name]
	if exist {
		// Eliminar el producto del inventario
		delete(inventory, name)
		fmt.Printf("Product '%s' deleted from inventory.\n", name)
	} else {
		fmt.Printf("The Product '%s' doesnt exist.\n", name)
	}
}

func ShowInventory(inventory map[string]Product) {
	fmt.Println("Inventory:")
	for name, product := range inventory {
		fmt.Printf("- Name: %s, Price: %.2f, Quantity: %d\n", name, product.price, product.quantity)
	}
}
