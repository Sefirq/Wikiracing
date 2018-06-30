package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	baseUrl := "https://en.wikipedia.org/wiki/"
	p := new(WebsiteParser)
	fmt.Println("Hello! Type an article name.")
	for scanner.Scan() {
		articleName := scanner.Text()
		links, err := p.parse(baseUrl + articleName)
		if err == nil {
			fmt.Println(links)
			os.Exit(0)
		} else {
			fmt.Println(err)
			fmt.Println("There was an error. Do you want to pass another article name?")
			scanner.Scan()
			decision := scanner.Text()
			if decision == "n" || decision == "N" {
				fmt.Println("Bye bye!")
				os.Exit(0)
			} else {
				fmt.Println("Type an article name.")
			}
		}
	}

}
