package main

import (
  "log"
  "net"
  "strconv"
  "time"
  "os"
  // "fmt"
  // "flag"
)

var ip = os.Args[1]
var minPort = 1
var maxPort = 65535

func main() {
  activeThreads := 0
  doneChannel := make(chan bool)
  // -help --help
  // flag.Usage = func() {
  //   fmt.Println("Useage-1 : go run tcp_port_scan.go ip_address")
  // go build tcp_port_scan.go
  //   fmt.Println("Useage-2 : ./tcp_port_scan ip_address")
  //   flag.PrintDefaults()
  // }
  // flag.Parse()

  for port := minPort; port<=maxPort; port++ {
    go tcpConnection(ip, port, doneChannel) // go threads
    activeThreads++
  }
  // wait for all threads to finish
  for activeThreads > 0 {
    <-doneChannel
    activeThreads--
  }
}

func tcpConnection(ip string, port int, doneChannel chan bool) {
  _, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), time.Second*10)
  if err == nil {
    log.Printf("Host %s has open port: %d\n", ip, port)
  }
  doneChannel <- true
}
