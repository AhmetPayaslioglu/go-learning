package main

import (
	"fmt"
	"net/http"
	"os"
	"golang.org/x/net/html"
)

func main() {
	baseURL := os.Args[1]
	response, err := http.Get(baseURL)
	checkError(err)

	defer response.Body.Close()
	doc, err := html.Parse(response.Body)
	checkError(err)
	// Recursively visit nodes in the parse tree
	var f func(*html.Node)
	f = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "a" {
					for _, a := range n.Attr {
							if a.Key == "href" {
									fmt.Println(a.Val)
									break
							}
					}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
					f(c)
			}
	}
	f(doc)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
