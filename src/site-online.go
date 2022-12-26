package main

import "fmt"

func main() {
	var name string = "Jose Valter"
	var version float32 = 1.1
	var option int

	fmt.Println("Hello", name)
	fmt.Println("This program is in the version", version)

	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("3 - Exit")

	fmt.Print("Type the option: ")
	fmt.Scan(&option)

	if option == 1 {
		fmt.Println("Monitoring...")
	} else if option == 2 {
		fmt.Println("Displaying logs...")
	} else if option == 3 {
		fmt.Println("Leaving the program")
	} else {
		fmt.Println("I don't know this command")
	}
}
