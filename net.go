package mock

import (
	"errors"
	"net"
	"time"
)

//shouldMockFail calls and return error?
var shouldMockFail = false

//ReturnError makes function calls return error if possible
func ReturnError(flag bool) {
	shouldMockFail = flag
}

//ConnMock is mock net.Conn struct for testing
type ConnMock struct {
}

//Read implement for net.Conn interface
func (m ConnMock) Read(b []byte) (n int, err error) {
	if shouldMockFail {
		err = errors.New("Dummy Error")
	}
	return
}

//Write implement for net.Conn interface
func (m ConnMock) Write(b []byte) (n int, err error) {
	if shouldMockFail {
		err = errors.New("Dummy Error")
	}
	return
}

//Close implement for net.Conn interface
func (m ConnMock) Close() error {
	if shouldMockFail {
		return errors.New("Dummy Error")
	}
	return nil
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
	if shouldMockFail {
		return errors.New("Dummy Error")
	}
	return nil
}

//SetReadDeadline implement for net.Conn interface
func (m ConnMock) SetReadDeadline(t time.Time) error {
	if shouldMockFail {
		return errors.New("Dummy Error")
	}
	return nil
}

//SetWriteDeadline implement for net.Conn interface
func (m ConnMock) SetWriteDeadline(t time.Time) error {
	if shouldMockFail {
		return errors.New("Dummy Error")
	}
	return nil
}
