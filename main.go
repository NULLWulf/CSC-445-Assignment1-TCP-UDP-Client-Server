package main

import (
	"flag"
	"fmt"
	"os"
	//"os/signal"
)

func main() {

	op := flag.String("mode", "", "Application Mode: Server (s) or Client (c)")
	addr := flag.String("addr", "", "Address to connect to")
	prtc := flag.String("prot", "", "Protocol to use: TCP (t) or UDP (u)")
	flag.Parse()

	switch *prtc {
	case "t":
		if *op == "s" {
			fmt.Println("Starting TCP server")
			tcpServer()
		} else if *op == "c" {
			fmt.Println("Starting TCP client")
			fmt.Println("Measuring RTT and ThroughPut in TCP Program Starting...")
			measureRtt(*addr)
			startThroughPutMeasurement(*addr, false)
		} else {
			fmt.Println("Invalid operation...Exiting Pleas see help (-h) for more info")
			os.Exit(1)
		}
	case "u":
		if *op == "s" {
			fmt.Println("Starting UDP Listener")
			//tcpServer()
		} else if *op == "c" {
			fmt.Println("Starting UDP client")
			fmt.Println("Measuring RTT and Throughput in UDP Program Starting...")
			measureRtt(*addr)
			startThroughPutMeasurement(*addr, true)
		} else {
			fmt.Println("Invalid operation...Exiting Pleas see help (-h) for more info")
			os.Exit(2)
		}
		break
	default:
		fmt.Println("Invalid protocol...Exiting Pleas see help (-h) for more info")
		os.Exit(1)
	}
	os.Exit(-5)
}
