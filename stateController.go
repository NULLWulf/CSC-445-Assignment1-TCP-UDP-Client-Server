package main

func runProgram() {
	switch Mode {
	case "server":
		switch Protocol {
		case "tcp":
			startTCPServer()
			break
		case "udp":
			startUDPServer()
			break
		}
		break
	case "client":
		measureRtt()
		startThroughPutMeasurement()
		break
	}
}
