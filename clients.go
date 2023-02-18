package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"net"
	"reflect"
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
		msg = XOREncode(msg)

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
		reply = XORDecode(reply)

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

	msg = XOREncode(msg)

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

func AssertEqual(data []byte) {
	encoded := XOREncode(data)
	decoded := XORDecode(encoded)
	if !reflect.DeepEqual(data, decoded) {
		fmt.Printf("Assertion Error: expected %v but got %v", data, decoded)
	} else {
		fmt.Println("Assertion Passed!")
	}
}

func XOREncode(data []byte) []byte {
	var result []byte
	for i := 0; i < len(data); i += 8 {
		var block uint64
		for j := i; j < i+8 && j < len(data); j++ {
			block |= uint64(data[j]) << ((j - i) * 8)
		}
		for j := 0; j < 8; j++ {
			result = append(result, byte((block>>(j*8))&0xff))
		}
	}
	for i := 8; i < len(result); i += 8 {
		for j := 0; j < 8; j++ {
			result[i+j] ^= result[i+j-8]
		}
	}
	return result
}

func XORDecode(data []byte) []byte {
	var result []byte
	for i := len(data) - 8; i >= 0; i -= 8 {
		var block uint64
		for j := 0; j < 8; j++ {
			block |= uint64(data[i+j]) << (j * 8)
		}
		for j := 0; j < 8; j++ {
			result = append(result, byte((block>>(j*8))&0xff))
		}
		if i > 0 {
			for j := 0; j < 8; j++ {
				result[i+j-8] ^= data[i-8+j]
			}
		}
	}
	return result
}
