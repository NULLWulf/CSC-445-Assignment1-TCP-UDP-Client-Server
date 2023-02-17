package main

import (
	"fmt"
	"net"
)

//func measureThroughputUDP(addr string, msgSize int, numMsgs int) (float64, error) {
//	conn, err := net.Dial("udp", addr)
//	if err != nil {
//		return 0, fmt.Errorf("error connecting to server: %s", err)
//	}
//	defer func(conn net.Conn) {
//		err := conn.Close()
//		if err != nil {
//
//		}
//	}(conn)
//
//	// create message and acknowledgement
//	msg := make([]byte, msgSize)
//	ack := make([]byte, 8)
//	binary.LittleEndian.PutUint64(ack, 1)
//
//	// encode message
//	xorEncode(msg)
//
//	// send messages and measure throughput
//	start := time.Now()
//	for i := 0; i < numMsgs; i++ {
//		_, err = conn.Write(msg)
//		if err != nil {
//			return 0, fmt.Errorf("error sending message: %s", err)
//		}
//		_, err = conn.Read(ack)
//		if err != nil {
//			return 0, fmt.Errorf("error receiving acknowledgement: %s", err)
//		}
//	}
//	elapsed := time.Since(start)
//
//	// calculate throughput
//	totalBytes := float64(msgSize * numMsgs)
//	throughput := totalBytes / elapsed.Seconds() * 8
//
//	return throughput, nil
//}

func handleConnectionUDP(conn *net.UDPConn) {
	// create message buffer
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
		xorDecode(msg)
		fmt.Printf("Received message from %s: %s\n", raddr, msg)
		// send acknowledgement
		_, err = conn.WriteToUDP([]byte("ACK"), raddr)
		if err != nil {
			fmt.Println("Error sending acknowledgement:", err)
			continue
		}
	}
}

func startUDPServer() {
	addr := &net.UDPAddr{IP: net.IPv4zero, Port: 12530}
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
