package gonetprobe

import (
	"fmt"
	"net"
	"time"
)

type Dialer interface {
	DialTimeout(network, address string, timeout time.Duration) (net.Conn, error)
}

type MockDialer struct {
	shouldSucceed bool
}

func (md *MockDialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	if md.shouldSucceed {
		// Return a dummy net.Conn and a nil error
		return &net.TCPConn{}, nil
	} else {
		// Return a nil net.Conn and a dummy error
		return nil, fmt.Errorf("mock error")
	}
}
