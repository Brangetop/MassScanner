package main

func main() {

	ports := []int{22, 80, 443, 8080}

	scanIPforPorts("google.com", ports)
}
