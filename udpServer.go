package main

import (
	"fmt"
	"net"
)

func handleConnectionUDP(conn *net.UDPConn) {
	buf := make([]byte, 1024)

	for {
		// read message
		n, raddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading message:", err)
			continue
		}

		// decode message
		msg := buf[:n]

		// send acknowledgement
		_, err = conn.WriteToUDP(msg, raddr)
		if err != nil {
			fmt.Println("Error sending acknowledgement:", err)
			continue
		}
	}
}

func startUDPServer() {
	addr := &net.UDPAddr{IP: net.IPv4zero, Port: Port}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	fmt.Printf("Server listening on: %s\n", addr)

	handleConnectionUDP(conn)
}
