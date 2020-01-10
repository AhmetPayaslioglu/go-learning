package main

import (
  "log"
  "net"
  "strconv"
  "time"
  "os"
)

var ip = os.Args[1]
var minPort = 1
var maxPort = 65535

func main() {
  activeThreads := 0
  doneChannel := make(chan bool)
  for port := minPort; port<=maxPort; port++ {
    go banner(ip, port, doneChannel) // go threads
    activeThreads++
  }
  // wait for all threads to finish
  for activeThreads > 0 {
    <-doneChannel
    activeThreads--
  }
}

func banner(ip string, port int, doneChannel chan bool) {
  connection, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port), time.Second*10)
  if err != nil {
    doneChannel <- true
    return
  }
  // see if server offers anything to read
  buffer := make([]byte, 4096)
  connection.SetReadDeadline(time.Now().Add(time.Second*5))
  // set timeout
  numBytesRead, err := connection.Read(buffer)
  if err != nil {
    doneChannel <- true
    return
  }
  log.Printf("Banner from port %d\n%s\n", port, buffer[0:numBytesRead])
  doneChannel <- true
}
