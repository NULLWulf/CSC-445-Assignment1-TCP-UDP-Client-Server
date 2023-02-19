package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"log"
	"net"
	"reflect"
	"time"
)

func measureRtt() {
	sizes := []int{8, 32, 512, 1024}
	rttResults := make(map[int]time.Duration)

	for _, size := range sizes {
		msg := make([]byte, size)
		_, err := rand.Read(msg)
		conn, err := net.Dial(Protocol, Address)
		if err != nil {
			log.Println("Error connecting to server: ", err)
			return
		}
		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {
				log.Println("Error closing connection:", err)
			}
		}(conn)

		// send message and measure round-trip time
		start := time.Now().UnixNano()
		//log.Println(msg)

		_, err = conn.Write(msg)
		if err != nil {
			log.Println("Error sending message:", err)
			return
		}
		reply := make([]byte, size)
		_, err = conn.Read(reply)
		if err != nil {
			log.Println("Error receiving reply:", err)
			return
		}
		end := time.Now().UnixNano()
		log.Println(bytes.Equal(msg, reply))

		// store result
		rttResults[size] = time.Duration(end - start)
		time.Sleep(1 * time.Second)
	}

	// print results
	log.Println("Round-trip times:")
	for size, rtt := range rttResults {
		fmt.Printf("%d bytes: %v\n", size, rtt)
	}

}

// measureThroughput measures the throughput for a given message size and number of messages
// and returns the throughput in Mbps
func measureThroughput(msgSize int, numMsgs int) (float64, error) {
	conn, err := net.Dial(Protocol, Address)
	if err != nil {
		return 0, fmt.Errorf("error connecting to server: %s", err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("Error closing connection:", err)
		}
	}(conn)

	// create message and acknowledgement
	msg := make([]byte, msgSize)
	_, err = rand.Read(msg)
	ack := make([]byte, 8)

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

// startThroughPutMeasurement starts the throughput measurement for different message sizes and number of messages
// and calculates the throughput for each combination in Mbps
func startThroughPutMeasurement() {

	throughput, _ := measureThroughput(1024, 1024)
	fmt.Printf("Message size: %d bytes, number of messages: %d, throughput: %.2f Mbps\n", 1024, 1024, throughput*0.000001)

	throughput, _ = measureThroughput(512, 2048)
	fmt.Printf("Message size: %d bytes, number of messages: %d, throughput: %.2f Mbps\n", 512, 2048, throughput*0.000001)

	throughput, _ = measureThroughput(128, 8192)
	fmt.Printf("Message size: %d bytes, number of messages: %d, throughput: %.2f Mbps\n", 128, 8192, throughput*0.000001)
}

func AssertEqual(data []byte) {
	encoded := XOREncode(data)
	decoded := XORDecode(encoded)
	if !reflect.DeepEqual(data, decoded) {
		fmt.Printf("Assertion Error: expected %v but got %v", data, decoded)
	} else {
		log.Println("Assertion Passed!")
	}
}

func XOREncode(input []byte) []byte {
	output := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = input[i] ^ Key[i%len(Key)]
	}
	return output
}

func XORDecode(input []byte) []byte {
	return XOREncode(input)
}
