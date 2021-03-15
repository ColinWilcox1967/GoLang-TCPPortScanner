package main

import (
	"fmt"
)

const (
	portScannerVersion = "1.0"
)

type scanSettings struct {
	firstPort, lastPort int

}

var settings scanSettings


func showTitle () {
	fmt.Printf ("TCP/IP Port Scanner (version %s)\n", portScannerVersion)
}

func getCommandLineArguments () scanSettings {

	return scanSettings{}
}

func main () {
	showTitle ()

	settings = getCommandLineArguments ()
}