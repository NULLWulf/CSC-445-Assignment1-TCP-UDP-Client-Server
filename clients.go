package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func measureRtt() {
	sizes := []int{8, 32, 512, 1024}
	rttResults := make(map[int]time.Duration)

	for _, size := range sizes {
		msg := make([]byte, size)
		conn, err := net.Dial(Protocol, Address)

		if err != nil {
			fmt.Println("Error connecting to server:", err)
			return
		}
		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {

			}
		}(conn)

		// encode message
		xorEncode(msg)

		// send message and measure round-trip time
		start := time.Now().UnixNano()
		_, err = conn.Write(msg)
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
		reply := make([]byte, size)
		_, err = conn.Read(reply)
		if err != nil {
			fmt.Println("Error receiving reply:", err)
			return
		}
		end := time.Now().UnixNano()

		// decode reply
		xorDecode(reply)

		// validate reply
		if !bytes.Equal(msg, reply) {
			fmt.Printf("Validation failed for size %d\n", size)
		}

		// store result
		rttResults[size] = time.Duration(end - start)
	}

	// print results
	fmt.Println("Round-trip times:")
	for size, rtt := range rttResults {
		fmt.Printf("%d bytes: %v\n", size, rtt)
	}

}

func measureThroughput(msgSize int, numMsgs int) (float64, error) {
	conn, err := net.Dial(Protocol, Address)
	if err != nil {
		return 0, fmt.Errorf("error connecting to server: %s", err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error closing connection:", err)
		}
	}(conn)

	// create message and acknowledgement
	msg := make([]byte, msgSize)
	_, err = rand.Read(msg)
	ack := make([]byte, 8)
	binary.LittleEndian.PutUint64(ack, 1)

	// encode message
	msg = xorEncode(msg)

	// send messages and measure throughput
	start := time.Now().UnixNano()
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
	end := time.Now().UnixNano()
	elapsed := time.Duration(end - start)

	// calculate throughput
	totalBytes := float64(msgSize * numMsgs)
	throughput := totalBytes / elapsed.Seconds() * 8

	return throughput, nil
}

func startThroughPutMeasurement() {
	msgSizes := []int{1024, 512, 128}
	numMsgs := []int{1024, 2048, 8192}

	for _, size := range msgSizes {
		for _, num := range numMsgs {
			throughput, err := measureThroughput(size, num)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Message size: %d bytes, number of messages: %d, throughput: %.2f Mbps\n", size, num, throughput/1000000)
		}
	}
}

func xorEncode(msg []byte) []byte {
	key := uint64(0x1234567890ABCDEF)
	for i := 0; i < len(msg); i += 8 {
		block := msg[i : i+8]
		value := key ^ binary.LittleEndian.Uint64(block)
		binary.LittleEndian.PutUint64(block, value)
	}

	return msg
}

func xorDecode(msg []byte) []byte {
	return xorEncode(msg)
}
