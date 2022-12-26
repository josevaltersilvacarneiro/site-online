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

	fmt.Scanf("%d", &option)
}
