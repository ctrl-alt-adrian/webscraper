package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	greetUser()

	// get url and format options  from command line
	websiteUrl, fetchOption := getUserInput()

	url := addUrlPrefix(websiteUrl)

	// scrape site
	scraper(url, fetchOption)

}

func greetUser() {
	fmt.Println("Welcome. Let's get started")
}

func getUserInput() (string, string) {
	var websiteUrl string
	var fetchOption string

	fmt.Println("Enter the website url that you want to scrape")
	fmt.Scanln(&websiteUrl)

	fmt.Println("What do you want to retrieve? Just the links, images, or all of the above")
	fmt.Scanln(&fetchOption)

	return websiteUrl, fetchOption
}

func addUrlPrefix(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	return url
}

func saveToFile(data []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	for _, entry := range data {
		_, err := file.WriteString(entry + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
