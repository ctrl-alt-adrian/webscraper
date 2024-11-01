package main

import (
	"fmt"
)

func main() {

	greetUser()

	// get url and format options  from command line
	websiteUrl, fetchOption := getUserInput()

	// scrape site
	scraper(websiteUrl, fetchOption)

}

func greetUser() {
	fmt.Println("Welcome. Let's get started")
}

func getUserInput() (string, string) {
	var websiteUrl string
	var fetchOption string

	fmt.Println("Enter the website url that you want to scrape")
	fmt.Scanln(&websiteUrl)

	fmt.Println("What do you want to retrieve? Just the links, or images")
	fmt.Scanln(&fetchOption)

	return websiteUrl, fetchOption
}
