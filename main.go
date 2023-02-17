package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	//"os/signal"
)

func main() {

	parseProgramArguments()
	SetupCloseHandler()

	switch Protocol {
	case "t":
		if Mode == "s" {
			fmt.Println("Starting TCP server")
			tcpServer()
		} else if Mode == "c" {
			fmt.Println("Starting TCP client")
			fmt.Println("Measuring RTT and ThroughPut in TCP Program Starting...")
			measureRtt(Address, false)
			startThroughPutMeasurement()
		} else {
			fmt.Println("Invalid operation...Exiting Pleas see help (-h) for more info")
			os.Exit(1)
		}
	case "u":
		if Mode == "s" {
			fmt.Println("Starting UDP Listener")
			startUDPServer()
		} else if Mode == "c" {
			fmt.Println("Starting UDP client")
			fmt.Println("Measuring RTT and Throughput in UDP Program Starting...")
			measureRtt(Address, true)
			startThroughPutMeasurement()
		} else {
			fmt.Println("Invalid operation...Exiting Pleas see help (-h) for more info")
			os.Exit(2)
		}
		break
	default:
		fmt.Println("Invalid Protocol...Exiting Pleas see help (-h) for more info")
		os.Exit(1)
	}
	os.Exit(-5)
}

func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
}
