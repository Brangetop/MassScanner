package main

import (
	"fmt"
	"strconv"
)

func main() {

	ports := []int{22, 80, 443, 8080}

	scanRange("87.251.74.", ports)
}

func scanRange(r string, p []int) {
	for i := 0; i <= 255; i++ {
		ip := r + strconv.Itoa(i)

		fmt.Println("scanning", ip)
		scanIPforPorts(ip, p)

	}
}
