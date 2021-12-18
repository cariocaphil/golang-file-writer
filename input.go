package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/cariocaphil/golang-file-writer/info"
)

func newProduct(id string, title string, description string, price float64) *product {
	return &product{id, title, description, price}
}

func getProductInput() *product {
	info.PrintWelcome()

	reader := bufio.NewReader(os.Stdin)

	idInput := readUserInput(reader, "Product ID: ")
	titleInput := readUserInput(reader, "Product Title: ")
	descriptionInput := readUserInput(reader, "Product Description: ")
	priceInput := readUserInput(reader, "Product Price: ")
	priceValue, _ := strconv.ParseFloat(priceInput, 64)

	product := newProduct(idInput, titleInput, descriptionInput, priceValue)

	return product
}

func readUserInput(reader *bufio.Reader, promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')

	if runtime.GOOS == "windows" {
		userInput = strings.Replace(userInput, "\r\n", "", -1)
	} else {
		userInput = strings.Replace(userInput, "\n", "", -1)
	}

	return userInput
}
