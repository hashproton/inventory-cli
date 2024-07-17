package main

import (
	"fmt"
	"go-api/models"
)

var products = []models.Product{}

func main() {
	seed()

	for {
		chooseOption()
	}
}

func chooseOption() {
	println("1 - List all products")
	println("2 - Add a new product")

	var option int

	fmt.Print("Choose an option: ")
	if _, err := fmt.Scanf("%d", &option); err != nil {
		println("Invalid option")
		return
	}

	switch option {
	case 1:
		listProducts()
	case 2:
		addProduct()
	default:
		println("Invalid option")
	}

}

func listProducts() {
	for _, product := range products {
		println(product.Display())
	}
}

func addProduct() {
	var product models.Product

	print("Enter the product name: ")
	fmt.Scanln(&product.Name)

	print("Enter the product price: ")
	fmt.Scanln(&product.Price)

	product.ID = len(products) + 1

	products = append(products, product)

	println("Product added successfully")
}

func seed() {
	products = append(products, models.Product{ID: 1, Name: "Zelda Breath Wild", Price: 49.99})
	products = append(products, models.Product{ID: 2, Name: "Xbox Controller", Price: 62.20})
}
