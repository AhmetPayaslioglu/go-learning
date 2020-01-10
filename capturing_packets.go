/*
Please run as root 
go build capturing_packets.go
sudo ./capturing_packets
*/
package main

import (
  "fmt"
  "os"
  "time"
  "github.com/google/gopacket"
  "github.com/google/gopacket/layers"
  "github.com/google/gopacket/pcap"
  "github.com/google/gopacket/pcapgo"
)

var (
  device = "wlp2s0"
  snapshotLen int32 = 1024
  promiscuous = false
  err error
  timeout = -1 * time.Second
  handle *pcap.Handle
  packetCount = 0
)

func main() {
  // open output pcap file and write header
  f, _ := os.Create("network_traffic.pcap")
  w := pcapgo.NewWriter(f)
  w.WriteFileHeader(uint32(snapshotLen), layers.LinkTypeEthernet)
  defer f.Close()
  // open the device capturing
  handle, err := pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
  if err != nil {
    fmt.Println("Error opening device %s: %v", device, err)
    os.Exit(1)
  }
  defer handle.Close()
  // start processing packets
  packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
  for packet := range packetSource.Packets() {
    // process packet here
    fmt.Println(packet)
    w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
    packetCount++
    // only capture 50 and stop
    if packetCount > 50 {
      break
    }
  }
}
