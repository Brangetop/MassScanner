package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func scanIPforPorts(ip string, ports []int) {
	timeout := 500 * time.Millisecond

	var wg sync.WaitGroup

	for _, port := range ports {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			address := net.JoinHostPort(ip, strconv.Itoa(p))

			conn, err := net.DialTimeout("tcp", address, timeout)
			if err != nil {
				return
			}

			fmt.Println("port is open: ", p)

			conn.Close()
		}(port)
	}

	wg.Wait()
}
