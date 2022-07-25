package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var host *string

func init() {
	host = flag.String("host", "scanme.nmap.org", "-h")
}

func main() {
	flag.Parse()
	fmt.Printf("Scanning %s\n", *host)
	var wg sync.WaitGroup
	for i := 0; i < 65535; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			port := i
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *host, port))
			if err != nil {
				return
			}
			conn.Close()
			fmt.Println("Port", port, "is open")
		}(i)

	}
	wg.Wait()
}
