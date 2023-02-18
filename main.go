package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	//"os/signal"
)

var Key = []byte("d20a944716d86ef0")

func main() {

	AssertEqual([]byte("Hello World"))
	parseProgramArguments()
	SetupCloseHandler()
	runProgram()
	os.Exit(0)
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
