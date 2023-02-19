package main

import (
	"flag"
	"log"
	"os"
)

var (
	Address    string
	Mode       string
	Protocol   string
	Port       int
	Throughput bool
)

// parseProgramArguments parses the command line arguments and sets the global variables based on them
// if configuration is valid the program will continue, otherwise it will exit with an error code
func parseProgramArguments() {
	Key = []byte("d20a944716d86ef0")
	flag.StringVar(&Mode, "Mode", "", "Application mode: 'server' or client'.")
	flag.StringVar(&Address, "Address", "", "Remote address to connect to while in Client mode, this field is ignored when set in server mode.")
	flag.StringVar(&Protocol, "Protocol", "", "Application Protocol Mode: 'udp' or 'tcp'.")
	flag.IntVar(&Port, "Port", 7500, "Port the application will listen to while in server mode.")
	flag.BoolVar(&Throughput, "Throughput", false, "Specify throughput explicitly to send 8 byte acks back")
	flag.Parse()

	if Mode != "server" && Mode != "client" {
		log.Fatalf("Invalid Mode.  Must be either 'server' or 'client'.")
	}

	if Mode == "server" && Address != "" {
		log.Println("Warning: Address argument is ignored when application set to server mode.")
	}

	if Mode == "client" && Address == "" {
		log.Fatalf("Invalid Address.  Address must be specified for client mode.")
	}

	if Protocol != "tcp" && Protocol != "udp" {
		log.Print("Invalid Protocol.  Must be either 'tcp' or 'udp'.")
		os.Exit(-5)
	}

	if Throughput {
		log.Printf("Application seto throughput mode.  Will return 8 byte acks.")
	}
}
