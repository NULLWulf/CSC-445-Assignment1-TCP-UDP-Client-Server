package main

import (
	"log"
	"net"
)

func handleTCPConnection(conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error closing connection: %v\n", err)
		}
	}()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			switch err {
			case net.ErrClosed:
				log.Println("Connection closed by remote host")
			default:
				log.Printf("Error reading: %v\n", err)
			}
			return
		}

		log.Printf("Received message from %s: %d bytes", conn.RemoteAddr(), n)

		if _, err := conn.Write(buffer[:n]); err != nil {
			log.Printf("Error writing: %v\n", err)
			return
		}
	}
}

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
