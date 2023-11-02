package gonetprobe

import (
	"testing"
)

func TestTCPScan(t *testing.T) {
	// Test with a mock dialer that succeeds
	result := TCPScan(&MockDialer{shouldSucceed: true}, "localhost", 80)
	if !result {
		t.Errorf("Expected true, got %v", result)
	}

	// Test with a mock dialer that fails
	result = TCPScan(&MockDialer{shouldSucceed: false}, "localhost", 80)
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestScan(t *testing.T) {
	d := &MockDialer{shouldSucceed: true}
	hostname := "localhost"
	startingPort := 1
	endingPort := 1024
	numWorkers := 10

	results := Scan(d, "tcp", hostname, startingPort, endingPort, numWorkers)

	// Assert that the number of results matches the expected number of ports
	expectedNumResults := endingPort - startingPort + 1
	if len(results) != expectedNumResults {
		t.Errorf("Expected %d results, but got %d", expectedNumResults, len(results))
	}

	// Assert that all ports are within the expected range
	for _, result := range results {
		if result.Port < startingPort || result.Port > endingPort {
			t.Errorf("Port %d is out of range", result.Port)
		}
	}
}
