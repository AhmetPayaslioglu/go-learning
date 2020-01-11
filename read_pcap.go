/*
tcpdump to create a test file -> tcpdump -w test.pcap
go run read_pcap.go test.pcap
*/
package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
  "os"
)

var (
	pcapFile = os.Args[1]
	handle   *pcap.Handle
	err      error
)

func main() {
	// open file instead of device
	handle, err = pcap.OpenOffline(pcapFile)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	// loop through packets in file
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}
}
