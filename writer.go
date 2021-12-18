package main

type product struct {
	id          string
	title       string
	description string
	price       float64
}

func main() {
	createdProduct := getProductInput()
	printData(createdProduct)
	storeData(createdProduct)
}
