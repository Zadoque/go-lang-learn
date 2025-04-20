package main

import "fmt"

const LoginToken string = "sometoken" //Public

func main() {
	fmt.Println("Hello, world")
	var username string = "Zadoque"
	fmt.Println(username)
	fmt.Printf("My name is %T \n", username)

	var smallVal float64 = 255.4565315413
	fmt.Println(smallVal)

	var num int
	fmt.Println(num)

	//implicit type

	var website = "someurl"
	fmt.Println(website)

	//no var style

	numberOfStuff := 3000
	fmt.Println(numberOfStuff)

	// const var
	fmt.Println(LoginToken)

}
