#Network Port Scanner
A concurrent network port scanner written in Go.
##Prerequisites
Go 1.18 or higher
##Usage
Open main.go.
Edit target IP range and ports in main().
Run
##Configuration
maxRoutines: Controls concurrent scans (default: 100). Reduce on Windows if connection limits are reached. Increase on linux(1000 works well)