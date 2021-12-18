package main

import (
	"fmt"
	"os"
)

func printData(prod *product) {
	fmt.Printf("ID: %v\n", prod.id)
	fmt.Printf("Title: %v\n", prod.title)
	fmt.Printf("Description: %v\n", prod.description)
	fmt.Printf("Price: USD %.2f\n", prod.price)
}

func storeData(prod *product) {
	fileOutput, _ := os.Create(prod.id)

	content := fmt.Sprintf(
		`Id: %v
Title: %v
Description: %v
Price: USD %.2f
		`, prod.id, prod.title, prod.description, prod.price)

	fileOutput.WriteString(content)

	fileOutput.Close()
}
