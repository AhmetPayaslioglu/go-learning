package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func checkIfUrlExists(baseUrl, filePath string, doneChannel chan bool) {
	// Create URL object from raw string
	targetUrl, err := url.Parse(baseUrl)
	if err != nil {
		log.Println("Error parsing base URL. ", err)
	}
	// Set the part of the URL after the host name
	targetUrl.Path = filePath

	// Perform a HEAD only, checking status without
	response, err := http.Head(targetUrl.String())
	if err != nil {
		log.Println("Error fetching ", targetUrl.String())
	}

	// If server returns
	if response.StatusCode != 404 {
		fmt.Println(" Status Code :", response.StatusCode ," ", targetUrl.String())
	}

	// Signal completion so next thread can start
	doneChannel <- true
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println(os.Args[0] + " - Perform an HTTP HEAD request to a URL")
		fmt.Println("Usage: " + os.Args[0] +
			" <url> <wordlist_file> <maxThreads>")
		fmt.Println("Example: " + os.Args[0] +
			" wordlist.txt https://pausiber.xyz 10")
		os.Exit(1)
	}

	baseUrl := os.Args[1]
  wordlistFilename := os.Args[2]
	maxThreads, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal("Error converting maxThread value to integer. ", err)
	}

	activeThreads := 0
	doneChannel := make(chan bool)

	// open word list
	wordlistFile, err := os.Open(wordlistFilename)
	if err != nil {
		log.Fatal("Error opening wordlist file. ", err)
	}

	// Read each line and do an HTTP HEAD
	scanner := bufio.NewScanner(wordlistFile)
	for scanner.Scan() {
		go checkIfUrlExists(baseUrl, scanner.Text(), doneChannel)
		activeThreads++

		if activeThreads >= maxThreads {
			<-doneChannel
			activeThreads -= 1
		}
	}

	for activeThreads > 0 {
		<-doneChannel
		activeThreads -= 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading wordlist file. ", err)
	}
}
