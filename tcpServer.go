package main

import (
	"crypto/rand"
	"log"
	"net"
)

// handleTCPConnection Handles a single TCP or UDP connection connection
func handleTCPConnection(conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error closing connection: %v\n", err)
		}
	}()

	buf := make([]byte, 1024)
	// Continuously listens for incoming messages from connected client.
	for {
		n, err := conn.Read(buf)
		if err != nil {
			switch err {
			case net.ErrClosed:
				log.Println("Connection closed by remote host")
			default:
				log.Printf("Error reading: %v\n", err)
			}
			return
		}
		msg := buf[:n]
		msg = XORDecode(msg)
		msg = XOREncode(msg)
		log.Printf("Received message from %s: %d bytes", conn.RemoteAddr(), n)
		if Throughput {
			msg = make([]byte, 8)
			_, err = rand.Read(msg)
		}

		if _, err := conn.Write(buf); err != nil {
			log.Printf("Error writing: %v\n", err)
			return
		}
	}
}

// startTCP server, starts a TCP server and binds it to the globally
// specified port.
func startTCPServer() {
	addr := &net.TCPAddr{IP: net.IPv4zero, Port: Port}
	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalf("Error listening: %v\n", err)
	}
	defer func() {
		if err := ln.Close(); err != nil {
			log.Printf("Error closing the listener: %v\n", err)
		}
	}()

	log.Printf("Server listening on: %s\n", addr)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v\n", err)
			continue
		}
		go handleTCPConnection(conn)
	}
}
