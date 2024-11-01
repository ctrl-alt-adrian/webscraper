package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

var wg = sync.WaitGroup{}

func scraper(websiteUrl string, options string) {

	url := addUrlPrefix(websiteUrl)

	// make http request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making http request: ", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// get token
	tokenizer := html.NewTokenizer(resp.Body)

	fetchFromSelectedOption(tokenizer, options)
}

func addUrlPrefix(url string) string {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	return url
}

func fetchFromSelectedOption(tokenizer *html.Tokenizer, option string) {

	fmt.Printf("The token is %v, and the option selected is %v", tokenizer, option)

	var links []string
	var images []string

	wg.Add(1)
	if option == "links" {

		for {
			tokenType := tokenizer.Next()
			if tokenType == html.ErrorToken {
				break
			}

			if tokenType == html.StartTagToken {
				token := tokenizer.Token()

				fmt.Printf("The Token tag is %v", token.Data)
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
			saveToFile(links, "links.txt")
			fmt.Println("Links saved to links.text")
		}
	} else if option == "images" {
		for {
			tokenType := tokenizer.Next()

			if tokenType == html.ErrorToken {
				break
			}

			if tokenType == html.StartTagToken {
				token := tokenizer.Token()

				// check for "img" tag
				if token.Data == "img" {
					for _, attr := range token.Attr {
						// check for href attribute
						if attr.Key == "src" {
							// TODO: check if alt exists. if so replace file value with alt value

							// if attr.Key == "alt" {
							// 	images = append(images, attr.Val+" ("+attr.Val+")")
							// }
							images = append(images, attr.Val)
						}
						fmt.Println("images", images)
					}
				}
			}
			// TODO: save images inside of a directory instead of a text file
			saveToFile(images, "images.txt")
		}

	}
	wg.Wait()
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
	wg.Done()
	return nil
}
