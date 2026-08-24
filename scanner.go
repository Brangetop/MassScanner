package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func scanIPforPorts(ip string, ports []int) {
	timeout := 500 * time.Millisecond

	for _, port := range ports {
		address := net.JoinHostPort(ip, strconv.Itoa(port))

		conn, err := net.DialTimeout("tcp", address, timeout)
		if err != nil {
			continue
		}

		if err == nil {
			fmt.Println("port is open: ", port)

			conn.Close()
		}
	}
}
