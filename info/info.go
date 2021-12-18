package info

import (
	"fmt"
)

const mainTitle = "File Writer"
const separator = "--------------"
const productPrompt = "Please enter your product: "

func PrintWelcome() {
	fmt.Println(mainTitle)
	fmt.Println(separator)
}
