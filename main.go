package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	//"os/signal"
)

func main() {

	AssertEqual([]byte("Hello World"))
	parseProgramArguments()
	SetupCloseHandler()

	if Mode == "s" {

	}

	if Mode == "c" {

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
