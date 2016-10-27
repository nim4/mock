package mock

import (
	"net"
	"time"
)

//ConnMock is mock net.Conn struct for testing
type ConnMock struct {
	Error error
}

//Read implement for net.Conn interface
func (m ConnMock) Read(b []byte) (n int, err error) {
	if m.Error != nil {
		err = m.Error
	}
	return
}

//Write implement for net.Conn interface
func (m ConnMock) Write(b []byte) (n int, err error) {
	if m.Error != nil {
		err = m.Error
	}
	return
}

//Close implement for net.Conn interface
func (m ConnMock) Close() error {
	return m.Error
}

//LocalAddr implement for net.Conn interface
func (m ConnMock) LocalAddr() net.Addr {
	addr, _ := net.ResolveTCPAddr("tcp", ":80")
	return addr
}

//RemoteAddr implement for net.Conn interface
func (m ConnMock) RemoteAddr() net.Addr {
	return m.LocalAddr()
}

//SetDeadline implement for net.Conn interface
func (m ConnMock) SetDeadline(t time.Time) error {
	return m.Error
}

//SetReadDeadline implement for net.Conn interface
func (m ConnMock) SetReadDeadline(t time.Time) error {
	return m.Error
}

//SetWriteDeadline implement for net.Conn interface
func (m ConnMock) SetWriteDeadline(t time.Time) error {
	return m.Error
}
