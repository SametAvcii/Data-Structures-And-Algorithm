package main

import "fmt"

type Shirt struct {
	name  string
	price float64
}
type Pants struct {
	name  string
	price float64
}
type Coats struct {
	name  string
	price float64
}

func createProduct(productType string) interface{} {

	var product interface{}
	if productType == "shirt" {
		product = &Shirt{name: "Basic shirt", price: 25.0}
	} else if productType == "pants" {
		product = &Pants{name: "Basic pants", price: 35.0}
	} else if productType == "coats" {
		product = &Coats{name: "Basic coats", price: 55.0}
	}
	return product
}
func main() {
	product := createProduct("shirt")
	fmt.Println("product:", product)

	product = createProduct("pants")
	fmt.Println("product:", product)


	product = createProduct("coats")
	fmt.Println("product:", product)
}
