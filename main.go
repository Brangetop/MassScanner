package main

func main() {
	var ports []int
	ports = append(ports, 123, 80, 430)

	scanIPforPorts("google.com", ports)
}
