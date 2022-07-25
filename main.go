package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var host *string
var from *int
var limit *int

func init() {
	host = flag.String("host", "scanme.nmap.org", "hostname to scan")
	from = flag.Int("from", 0, "from what port we will scan")
	limit = flag.Int("limit", 65535, "until what port we will scan")
}

func main() {
	flag.Parse()
	fmt.Printf("Scanning %s from port %d to %d\n", *host, *from, *limit)
	var wg sync.WaitGroup
	for i := *from; i < *limit; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
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
