package main

import (
	"fmt"
	"time"
	"net"
	"flag"
	"os"
)

const (
	portScannerVersion = "1.0"
)

type scanSettings struct {
	firstPort, lastPort int
	portTimeout    int
}

var settings scanSettings


func showTitle () {
	fmt.Printf ("TCP/IP Port Scanner (version %s)\n", portScannerVersion)
}

func showSyntax () {

}

func getCommandLineArguments () scanSettings {
	var ports string
	

	flag.StringVar(&ports, "ports", "", "Specifies the ports to be scanned.")
	portTimeout := flag.Int ("timeout", 1, "Time allowed for TCP reponse (in seconds).")
	if (*portTimeout <= 0) {
		*portTimeout = 1
	}

	settings.portTimeout = *portTimeout
	return settings
}

func tcpmultipleconnect (ip string, ports []string) map[string]string {
    // check emqx 1883, 8083 port

    results := make(map[string]string)
    for _, port := range ports {
        address := net.JoinHostPort(ip, port)
        // 3 second timeout
        conn, err := net.DialTimeout("tcp", address, 3*time.Second)
        if err != nil {
            results[port] = "failed"
            // todo log handler
        } else {
            if conn != nil {
                results[port] = "success"
                _ = conn.Close()
            } else {
                results[port] = "failed"
            }
        }
    }
    return results
}

func tcpconnect(host string, ports []string) {
    for _, port := range ports {
        timeout := time.Second
        conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
        if err != nil {
            fmt.Println("Connecting error:", err)
        }
        if conn != nil {
            defer conn.Close()
            fmt.Println("Opened", net.JoinHostPort(host, port))
        }
    }
}

func main () {
	showTitle ()
	if  len(os.Args[1:]) == 0 {
		showSyntax ()
		os.Exit(0)
	}

	settings = getCommandLineArguments ()
	
}