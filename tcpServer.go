package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing connection:", err)
		}
	}(conn)

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		if n == 0 {
			return
		}

		// decode message
		msg := buf[:n]
		msg = XORDecode(msg)
		msg = XOREncode(msg)

		fmt.Printf("Received message from %s: %s\n", conn.RemoteAddr(), msg)
		// echo message back to client
		_, err = conn.Write(msg)
		if err != nil {
			fmt.Println("Error writing:", err)
			return
		}
	}
}

func tcpServer() {
	addr := &net.TCPAddr{IP: net.IPv4zero, Port: Port}
	ln, err := net.ListenTCP("tcp", addr)

	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer func(ln net.Listener) {
		err := ln.Close()
		if err != nil {
			fmt.Println("Error closing the listener: ", err)
		}
	}(ln)

	fmt.Printf("Server listening on: %s\n", addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
