package main

import (
	"fmt"
	"log"
	"time"
  "os"
  "github.com/google/gopacket"
  "github.com/google/gopacket/pcap"
)

var (
	device            = "wlp2s0"
	snapshotLen int32 = 1024
	promiscuous       = false
	err         error
	timeout     = 30 * time.Second
	handle      *pcap.Handle
)

func main() {
	// open device
	handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// set filter
	var filter string = "TCP and port " + os.Args[1]
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Only capturing ", filter)

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Do something with a packet here.
		fmt.Println(packet)
	}

}
