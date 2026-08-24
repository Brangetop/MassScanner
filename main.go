package main

import (
	"fmt"
	"net/netip"
	"sync"
)

func main() {
	ports := []int{25565}

	scanRange("89.22.224.0/21", ports)
}

func scanRange(cidrStr string, p []int) {
	prefix, err := netip.ParsePrefix(cidrStr)
	if err != nil {
		fmt.Println("Wrong format:", cidrStr)
		return
	}

	var wg sync.WaitGroup

	maxRoutines := 100
	guard := make(chan struct{}, maxRoutines)

	currentIP := prefix.Masked().Addr()

	for prefix.Contains(currentIP) {
		guard <- struct{}{}
		wg.Add(1)

		go func(ip netip.Addr) {
			defer wg.Done()
			defer func() {
				<-guard
			}()

			ipStr := ip.String()

			scanIPforPorts(ipStr, p)
		}(currentIP)

		currentIP = currentIP.Next()
	}

	wg.Wait()
	close(guard)
	fmt.Println("Done")
}
