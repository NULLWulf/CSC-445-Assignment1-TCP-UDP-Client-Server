package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	//"os/signal"
)

func main() {
	parseProgramArguments()
	setupCloseHandler()
	runProgram()
	os.Exit(0)
}

// SetupCloseHandler handles the Ctrl+C event and exits gracefully
func setupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("\r- Ctrl+C pressed in Terminal.  Exiting gracefully")
		os.Exit(0)
	}()
}
