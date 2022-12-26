package main

import (
	"fmt"
	"os"
)

func main() {
	displayIntro()
	displayMenu()

	switch getOption() {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Displaying logs...")
	case 3:
		fmt.Println("Leaving the program")
		os.Exit(0)
	default:
		fmt.Println("I don't know this command")
		os.Exit(1)
	}
}

func displayIntro() {
	var name string = "Jose Valter"
	var version float32 = 1.1

	fmt.Println("Hello", name)
	fmt.Println("This program is in the version", version)
}

func displayMenu() {
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("3 - Exit")
}

func getOption() int {
	var option int

	fmt.Print("Type the option: ")
	fmt.Scan(&option)

	return option
}
