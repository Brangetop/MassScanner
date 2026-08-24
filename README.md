# Network Port Scanner

A concurrent network port scanner written in Go.

## Prerequisites

- Go 1.18 or higher

## Usage

1. Open `main.go`.
2. Edit the target IP range and ports in the `main()` function.
3. Run the scanner:

   ```bash
   go run main.go
## Configuration
maxRoutines controls the number of concurrent scans.

  ```go


  maxRoutines := 100
  Default: 100
```
Reduce the value on Windows if connection limits are reached.
Increase the value on Linux. A value of 1000 works well.
