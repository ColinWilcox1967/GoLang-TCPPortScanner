package main

import (
	"fmt"
	"time"
	"net"
	"flag"
    "strings"
	"os"
    "strconv"
)

const (
	portScannerVersion = "1.0"
)

var (
    portTimeout int
    portList []int
    host string
)

func showTitle () {
	fmt.Printf ("TCP/IP Port Scanner (version %s)\n", portScannerVersion)
}

func showSyntax () {
    fmt.Println ("Syntax: PORTSCAN [-port=<port list>] [-timeout=<period in seconds>] [-host=<host name or IP>]")
}

func getCommandLineArguments () int {
	var ports string

    // host
    flag.StringVar (&host, "host", "golang.org:8080", "Specifies host URL or IP")

    //timeout
	portTimeout := flag.Int ("timeout", 1, "Time allowed for TCP reponse (in seconds).")

   	if (*portTimeout <= 0) {
		*portTimeout = 1
	}

    fmt.Printf ("Host:'%s' (Timeout %ds)\n\n", host, *portTimeout)

    //port
	flag.StringVar(&ports, "port", "", "Specifies the ports to be scanned.")

    allPorts := strings.Split(ports, ",")

    for _, portstr := range allPorts {
        if len(portstr) > 0 {
            fmt.Printf ("'%s' ", portstr)
            port, err := strconv.Atoi(portstr)
            if err == nil {
                portList = append(portList, int(port))
            } else {
                fmt.Printf ("Invalid port specified ('%d').\n", port)
            }
        }
    }

    portList = append(portList, 30000)
    portList = append(portList, 30001)
    return len(portList)
}

func tcpConnect(host string, port string) {
    conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), time.Duration(portTimeout))
    if err != nil {
        fmt.Println("Connecting error:", err)
    }
    if conn != nil {
        defer conn.Close()
        fmt.Println("Opened", net.JoinHostPort(host, port))
    }
 }

func main () {
	showTitle ()
	if  len(os.Args[1:]) == 3 {
		showSyntax ()
		os.Exit(0)
	}
 
	countPorts := getCommandLineArguments ()

    if countPorts == 0 {
        fmt.Println ("Error: No valid ports specified.")
        os.Exit(-1)
    }
    // now iterate over all defined ports
    for port,_ := range portList {
        tcpConnect (host, fmt.Sprintf("%d", port))
    }
}