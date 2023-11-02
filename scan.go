package gonetprobe

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// ScanPort scans a port
func ScanPort(dialer Dialer, protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := dialer.DialTimeout(protocol, address, 60*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

// InitialScan scans the first 1024 ports
func InitialScan(dialer Dialer, hostname string) {
	var wg sync.WaitGroup
	for i := 0; i <= 1024; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			result := ScanPort(dialer, "tcp", hostname, port)
			if result {
				fmt.Println("tcp port", port, "is open")
			}
		}(i)
	}
	wg.Wait()
}

// WideScan scans the first 49152 ports
func WideScan(dialer Dialer, hostname string) {
	var wg sync.WaitGroup
	for i := 0; i <= 49152; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			result := ScanPort(dialer, "tcp", hostname, port)
			if result {
				fmt.Println("tcp port", port, "is open")
			}
		}(i)
	}
	wg.Wait()
}
