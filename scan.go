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

func Scan(dialer Dialer, protocol, hostname string, startingPort, endingPort, numWorkers int) []ScanResult {
	var wg sync.WaitGroup
	ports := make(chan int)
	results := make(chan ScanResult)

	// Start the workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1) // Increment WaitGroup counter for each worker
		go func() {
			defer wg.Done() // Call Done() when the worker finishes
			for port := range ports {
				result := ScanPort(dialer, protocol, hostname, port)
				results <- ScanResult{Port: port, Open: result}
			}
		}()
	}

	// Send ports to be scanned
	go func() {
		for i := startingPort; i <= endingPort; i++ {
			ports <- i
		}
		close(ports) // Close the channel after all ports have been sent
	}()

	go func() {
		wg.Wait()      // Wait for all workers to finish
		close(results) // Close the results channel after all workers have finished
	}()

	// Collect the results
	var scanResults []ScanResult
	for result := range results {
		scanResults = append(scanResults, result)
	}

	return scanResults
}

// InitialScan scans the first 1024 ports
func InitialTCPScan(dialer Dialer, hostname string, numWorkers int) []ScanResult {
	return Scan(dialer, "tcp", hostname, 1, 1024, numWorkers)
}

// WideScan scans the first 49152 ports
func WideTCPScan(dialer Dialer, hostname string, numWorkers int) []ScanResult {
	return Scan(dialer, "tcp", hostname, 1, 49152, numWorkers)
}
