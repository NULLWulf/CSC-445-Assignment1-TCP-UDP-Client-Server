package main

func runProgram() {

	switch Mode {
	case "s":
		switch Protocol {
		case "t":
			startTCPServer()
			break
		case "u":
			startUDPServer()
			break
		}
		break
	case "c":
		measureRtt()
		startThroughPutMeasurement()
		break
	}
}
