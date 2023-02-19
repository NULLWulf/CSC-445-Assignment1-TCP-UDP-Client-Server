package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net"
)

// handleConnectipnUDP handles a single udp "connection"
func handleConnectionUDP(conn *net.UDPConn) {
	buf := make([]byte, 1024)

	for {
		// read message
		n, raddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		// decode message
		msg := buf[:n]
		//log.Printf("Received message from %s: %d bytes", conn.RemoteAddr(), n)
		if Throughput {
			msg = make([]byte, 8)
			_, err = rand.Read(msg)
		}

		// send acknowledgement
		_, err = conn.WriteToUDP(msg, raddr)
		if err != nil {
			log.Println("Error sending acknowledgement:", err)
			continue
		}
	}
}

// startUDPServer starts UDP server instance and binds to globally defined port.
func startUDPServer() {
	addr := &net.UDPAddr{IP: net.IPv4zero, Port: Port}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Println("Error starting server:", err)
		return
	}
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			log.Println("Error closing the connection:", err)
		}
	}(conn)

	fmt.Printf("Server listening on: %s\n", addr)

	handleConnectionUDP(conn)
}
