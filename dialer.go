package gonetprobe

import (
	"errors"
	"net"
	"time"
)

type Dialer interface {
	DialTimeout(network, address string, timeout time.Duration) (net.Conn, error)
}

type MockDialer struct {
	shouldSucceed bool
}

func (d *MockDialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	if d.shouldSucceed {
		// Return a successful connection
		return &MockConn{}, nil
	} else {
		// Return an error to simulate failure
		return nil, errors.New("connection failed")
	}
}
