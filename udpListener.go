package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func measureThroughputUDP(addr string, msgSize int, numMsgs int) (float64, error) {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		return 0, fmt.Errorf("error connecting to server: %s", err)
	}
	defer conn.Close()

	// create message and acknowledgement
	msg := make([]byte, msgSize)
	ack := make([]byte, 8)
	binary.LittleEndian.PutUint64(ack, 1)

	// encode message
	xorEncode(msg)

	// send messages and measure throughput
	start := time.Now()
	for i := 0; i < numMsgs; i++ {
		_, err = conn.Write(msg)
		if err != nil {
			return 0, fmt.Errorf("error sending message: %s", err)
		}
		_, err = conn.Read(ack)
		if err != nil {
			return 0, fmt.Errorf("error receiving acknowledgement: %s", err)
		}
	}
	elapsed := time.Since(start)

	// calculate throughput
	totalBytes := float64(msgSize * numMsgs)
	throughput := totalBytes / elapsed.Seconds() * 8

	return throughput, nil
}
