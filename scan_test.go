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
