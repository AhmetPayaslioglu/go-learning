package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("No domain name argument provided")
	}
	arg := os.Args[1]

	mxRecords, err := net.LookupMX(arg)
	if err != nil {
		log.Fatal(err)
	}
	for _, mx := range mxRecords {
		fmt.Println(mx.Host, mx.Pref)
	}
}
