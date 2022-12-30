package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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
			showLogs()
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

func showLogs() {
	/*
		This procedure opens the file `log.txt`
		and reads its content. The content
		is received by the variable `file` and
		the error is received by the variable
		`err`. If there is an error, this pro-
		cedure shows it. At the end, it prints
		the logs.
	*/

	fmt.Println("Displaying logs...")

	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("There was an error!", err)
	}

	fmt.Println(string(file))
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
		logRecord(site, true)
	} else {
		fmt.Println("The site", site, "is in trouble. Status:",
			response.StatusCode)
		logRecord(site, false)
	}
}

func logRecord(site string, status bool) {
	/*
		This procedure opens the file `log.txt` in
		add mode. If there is an error, it will show
		the error. At the end, it stores the status
		of the site in the `log.txt` file.
	*/

	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("There was an error!", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 at 15/04/05") + " - " + site +
		" - Online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}
