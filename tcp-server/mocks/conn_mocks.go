package mocks

import (
	"net"
	"time"

	"github.com/stretchr/testify/mock"
)

type ConnectionMock struct {
	mock.Mock
}

func (conn *ConnectionMock) Read(b []byte) (n int, err error) {
	args := conn.Called(b)
	return args.Int(0), args.Error(1)
}

func (conn *ConnectionMock) Write(b []byte) (n int, err error) {
	args := conn.Called(b)
	return args.Int(0), args.Error(1)
}

func (conn *ConnectionMock) Close() error {
	return nil
}

func (conn *ConnectionMock) LocalAddr() net.Addr {
	return nil
}

func (conn *ConnectionMock) RemoteAddr() net.Addr               { return &AddrMock{} }
func (conn *ConnectionMock) SetDeadline(t time.Time) error      { return nil }
func (conn *ConnectionMock) SetReadDeadline(t time.Time) error  { return nil }
func (conn *ConnectionMock) SetWriteDeadline(t time.Time) error { return nil }

type AddrMock struct {
	mock.Mock
}

func (addr *AddrMock) Network() string { return "" }
func (addr *AddrMock) String() string  { return "127.0.0.1:56145" }
