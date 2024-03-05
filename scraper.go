package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	// get url from command line
	url := os.Args[1]

	// make http request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making http request: ", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// parse html
	tokenizer := html.NewTokenizer(resp.Body)

	// extract links from html
	var links []string
	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			break
		}

		if tokenType == html.StartTagToken {
			token := tokenizer.Token()

			// check for "a" tag
			if token.Data == "a" {
				for _, attr := range token.Attr {
					// check for href attribute
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			}
		}

	}

	// save links to a txt file
	saveToFile(links, "links.txt")
	fmt.Println("Data saved to links.text")
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
