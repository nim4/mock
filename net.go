package mock

import (
	"io"
	"net"
	"time"
)

//ConnMock is mock net.Conn struct for testing
type ConnMock struct {
	Buffer [][]byte
	Error  error
	i      int
}

//Read implement for net.Conn interface
func (m ConnMock) Read(b []byte) (int, error) {
	if m.Error != nil {
		return 0, m.Error
	}

	if m.Buffer != nil {
		if m.i < len(m.Buffer) {
			b = m.Buffer[m.i]
			m.i++
			return len(b), nil
		}
		return 0, io.EOF
	}

	return 0, nil
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
