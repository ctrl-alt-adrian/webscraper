package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func scraper(url string, options string) {

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

func fetchFromSelectedOption(tokenizer *html.Tokenizer, option string) {
	// will need to pass in fetch options to retrieve only selected otion
	// will have to figure out how to save file to different directory location
	// will need to retrieve images and other information separately as an option

	fmt.Printf("The token is %v, and the option selected is %v", tokenizer, option)

	var links []string
	var images []string

	if option == "links" {
		for {
			tokenType := tokenizer.Next()
			fmt.Printf("The token type is %v", tokenType)
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
							// check if alt exists. if so replace file value with alt value
							if attr.Key == "alt" {
								images = append(images, attr.Val+" ("+attr.Val+")")
							}
							images = append(images, attr.Val)
						}
						fmt.Println("images", images)
					}
				}
			}
			saveToFile(images, "images.txt")
		}

	}
}
