package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const numberOfTimes = 5
const delay = 5

func main() {
	displayIntro()

	for {
		displayMenu()

		switch getOption() {
		case 1:
			startMonitoring()
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

func startMonitoring() {
	fmt.Println("Monitoring...")

	sites := getSites()

	for i := 0; i < numberOfTimes; i++ {
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(delay * time.Second)
	}
}

func getSites() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("There was an error!", err)
	}

	reader := bufio.NewReader(file)
	for {
		site, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("There was an error!", err)
		}

		sites = append(sites, strings.TrimSpace(site))
	}

	file.Close()

	return sites
}

func testSite(site string) {
	response, _ := http.Get(site)

	if response.StatusCode == 200 {
		fmt.Println("The site", site, "was loaded successfully")
	} else {
		fmt.Println("The site", site, "is in trouble. Status:",
			response.StatusCode)
	}
}
