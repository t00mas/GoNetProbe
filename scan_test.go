package gonetprobe

import (
	"testing"
)

func TestScanPort(t *testing.T) {
	// Test with a mock dialer that succeeds
	result := ScanPort(&MockDialer{shouldSucceed: true}, "tcp", "localhost", 80)
	if !result {
		t.Errorf("Expected true, got %v", result)
	}

	// Test with a mock dialer that fails
	result = ScanPort(&MockDialer{shouldSucceed: false}, "tcp", "localhost", 80)
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}
