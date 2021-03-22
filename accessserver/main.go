package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"./endpoints"

	mux "../github.com/gorilla/mux"
)

//
// API Endpoint handler functions
//

const (
	_DefaultPort   int    = 8080
	_ServerVersion string = "0.1"
)

var (
	_DefaultPortStr string
	helpFlag        *bool
	httpPort        int = _DefaultPort
	gmuxRouter          = mux.NewRouter().StrictSlash(true)
)

func showStartupMessage() {
	msg := fmt.Sprintf("Starting server on port %d\n", httpPort)
	fmt.Println(msg)
}

func startEngine() error {
	return nil

}

func handleRequests() {
	gmuxRouter.HandleFunc("/version", endpoints.GetEngineVersionEndpoint)
	gmuxRouter.HandleFunc("/port", endPoints.GetPortAddressEndpoint)

	portStr := fmt.Sprintf(":%d", httpPort)
	log.Fatal(http.ListenAndServe(portStr, gmuxRouter))
}

func handleCommandLineParameters() (int, error) {

	var str string
	flag.StringVar(&str, "port", _DefaultPortStr, "Port on which the user account server will run.")

	helpFlag = flag.Bool("help", false, "Help required by user.")

	// PArse the command line
	flag.Parse()

	// get the port number
	port, err := strconv.Atoi(str)

	return port, err
}

func showSyntax() {

	fmt.Printf("ASERVER [-HELP] | [-PORT = <port number>] \n\n")
	fmt.Printf("<port number>        - HTTP port number for server listener. Defaults to %d.\n", _DefaultPort)
}

func main() {

	var err error

	_DefaultPortStr = fmt.Sprintf("%d", _DefaultPort)

	httpPort, err = handleCommandLineParameters()

	if *helpFlag {
		showSyntax()
		os.Exit(0)
	}

	showStartupMessage()

	if err == nil {
		//		fmt.Printf("Starting User Accounts Server %s on Port %d ...\n", _ServerVersion, httpPort)
		handleRequests()
	}
}
