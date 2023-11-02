package gonetprobe

import (
	"net"
	"time"
)

type MockConn struct{}

func (c *MockConn) Read(b []byte) (n int, err error) {
	// Implement the Read method if needed for your specific test
	return 0, nil
}

func (c *MockConn) Write(b []byte) (n int, err error) {
	// Implement the Write method if needed for your specific test
	return 0, nil
}

func (c *MockConn) Close() error {
	// Implement the Close method if needed for your specific test
	return nil
}

func (c *MockConn) LocalAddr() net.Addr {
	// Implement the LocalAddr method if needed for your specific test
	return nil
}

func (c *MockConn) RemoteAddr() net.Addr {
	// Implement the RemoteAddr method if needed for your specific test
	return nil
}

func (c *MockConn) SetDeadline(t time.Time) error {
	// Implement the SetDeadline method if needed for your specific test
	return nil
}

func (c *MockConn) SetReadDeadline(t time.Time) error {
	// Implement the SetReadDeadline method if needed for your specific test
	return nil
}

func (c *MockConn) SetWriteDeadline(t time.Time) error {
	// Implement the SetWriteDeadline method if needed for your specific test
	return nil
}
