package gonetprobe

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type ScanResult struct {
	Port int
	Open bool
}

// ScanPort scans a port
func ScanPort(dialer Dialer, protocol, hostname string, port int) bool {
	switch protocol {
	case "tcp":
		return TCPScan(dialer, hostname, port)
	case "udp":
		fmt.Println("UDP scanning not implemented yet")
		return false
	default:
		return false
	}
}

func TCPScan(dialer Dialer, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := dialer.DialTimeout("tcp", address, 60*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func Scan(dialer Dialer, protocol, hostname string, startingPort, endingPort int) []ScanResult {
	var wg sync.WaitGroup
	ports := make(chan int)
	results := make(chan ScanResult)

	// Start the workers
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			for port := range ports {
				result := ScanPort(dialer, protocol, hostname, port)
				results <- ScanResult{Port: port, Open: result}
			}
		}()
	}

	// Send ports to be scanned
	for i := startingPort; i <= endingPort; i++ {
		wg.Add(1)
		ports <- i
	}

	close(ports)
	wg.Wait()

	// Collect the results
	var scanResults []ScanResult
	close(results)
	for result := range results {
		scanResults = append(scanResults, result)
	}

	return scanResults
}

// InitialScan scans the first 1024 ports
func InitialTCPScan(dialer Dialer, hostname string) []ScanResult {
	return Scan(dialer, "tcp", hostname, 1, 1024)
}

// WideScan scans the first 49152 ports
func WideTCPScan(dialer Dialer, hostname string) []ScanResult {
	return Scan(dialer, "tcp", hostname, 1, 49152)
}
