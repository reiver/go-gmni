package gemini

import (
	"crypto/tls"
	"io"
	"net"
)

// Conn represent a I/O connection.
//
// For example: a TCP connection.
type Conn interface {
	io.ReadWriteCloser

	RemoteAddr() net.Addr
	LocalAddr() net.Addr
}

// Dial makes an un-secure Naked client connection to the system's 'loopback address'
// (also known as "localhost" or 127.0.0.1).
//
// One would might use this function if one is trying to debug a Naked server.
// And one might create a Naked server because TLS encryption is being handled elsewhere.
//
// Normally one would instead use the `DialTLS` function.
func Dial() (Conn, error) {
	return DialTo("")
}

// DialTo makes an un-secure Nake client connection to the the address specified by `addr`.
//
// One would might use this function if one is trying to debug a Naked server.
// And one might create a Naked server because TLS encryption is being handled elsewhere.
//
// Normally one would instead use the `DialToTLS` function.
func DialTo(addr string) (Conn, error) {

	const network = "tcp"

	if "" == addr {
		addr = "127.0.0.1:"+string(DefaultNakedServerPort)
	}

	conn, err := net.Dial(network, addr)
	if nil != err {
		return nil, err
	}

	return conn, nil
}

// DialTLS makes a (secure) Gemini Protocol client connection to the system's 'loopback address'
// (also known as "localhost" or 127.0.0.1).
func DialTLS(tlsConfig *tls.Config) (Conn, error) {
	return DialToTLS("", tlsConfig)
}

// DialToTLS makes a (secure) Gemini Protocol client connection to the the address specified by `addr`.
func DialToTLS(addr string, tlsConfig *tls.Config) (Conn, error) {

	const network = "tcp"

	if "" == addr {
		addr = "127.0.0.1:"+string(DefaultServerPort)
	}

	conn, err := tls.Dial(network, addr, tlsConfig)
	if nil != err {
		return nil, err
	}

	return conn, nil
}
