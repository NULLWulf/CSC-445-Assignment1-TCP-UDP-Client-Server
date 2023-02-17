package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	Address  string
	Mode     string
	Protocol string
	Port     int
)

func parseProgramArguments() {
	// Get command line arguments and parse them
	flag.StringVar(&Mode, "Mode", "", "Application Mode: Server (s) or Client (c)")
	flag.StringVar(&Address, "Address", "", "Host Address to connect to or host Address to instantiate.")
	flag.StringVar(&Protocol, "Protocol", "", "Protocol to use: TCP (t) or UDP (u).")
	flag.IntVar(&Port, "Port", 0, "Port to listen on or seek on destination host.")
	flag.Parse()

	if Mode != "s" || Mode != "t" {
		fmt.Println("No Application Mode Detected. Setting Server By Default")
		Mode = "s"
	}
	if Address == "" {
		if Mode == "s" {
			fmt.Println("No Localhost Address Set for Server Mode, defaulting to 0.0.0.0")
			Address = "0.0.0.0"
		} else {
			fmt.Println("Error:  Application Set to Client Mode with Destination Host Address. Application Terminating")
			os.Exit(-4)
		}
	}
	if Protocol == "" {
		fmt.Println("No default Protocol specified.  Starting with TCP")
		Protocol = "t"
	}

	if Port == 0 {
		defPort := 6780
		if Mode == "s" {
			fmt.Println("App set to server Mode with no Port to listener, setting to default Port of : ", defPort)
		}
		Port = defPort
	}
}
