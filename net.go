package mock

import (
	"bytes"
	"io"
	"net"
	"time"
)

//ConnMock is mock net.Conn struct for testing
type ConnMock struct {
	Buffer [][]byte
	Error  error
}

//Read implement for net.Conn interface
func (m *ConnMock) Read(b []byte) (int, error) {
	if m.Error != nil {
		return 0, m.Error
	}

	if m.Buffer != nil {
		if len(m.Buffer) > 0 {
			reader := bytes.NewReader(m.Buffer[0])
			n, err := reader.Read(b)
			if n == len(m.Buffer[0]) {
				m.Buffer = m.Buffer[1:]
			} else {
				m.Buffer[0] = m.Buffer[0][n:]
			}
			return n, err
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
