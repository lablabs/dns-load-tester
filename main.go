package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func testDNS(urls []string, ch chan bool) {
	for {
		for _, url := range urls {
			lookup(url)
		}
	}
}

func lookup(url string) {
	ips, err := net.LookupIP(url)
	now := time.Now()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s - %s - Could not get IPs: %v\n", now, url, err)
	} else if strings.Contains(ips[0].String(), "::") {
		fmt.Fprintf(os.Stderr, "%s - %s - Got IPv6 address only: %v \n", now, url, ips)
	}

	// fmt.Println(ips)
}

func main() {
	threads := flag.Int("threads", 5, "how many lookups in parallel should run")
	flag.Parse()
	args := flag.Args()
	ch := make(chan bool)

	if len(args) < 1 {
		fmt.Println("No url specified")
		os.Exit(1)
	}

	for i := 0; i < *threads; i++ {
		go testDNS(args, ch)
	}
	<-ch
}
