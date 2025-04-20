package main

import (
	"bufio"
	"fmt"
	"os"
)

// how to get user input
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza: ")
	//comma ok || err ok
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of rating is %T \n", input)
}
