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
    defaultPortTimeout = 30
    defaultHostIP = "127.0.0.1"
    defaultPort = "21"
)

var (
    portTimeout int
    portList []int
    host string
)

func showTitle () {
	fmt.Printf ("TCP/IP Port Scanner (version %s)\n", portScannerVersion)
    fmt.Println("(c) Colin Wilcox 2021.")
}

func showSyntax () {
    fmt.Println ("Syntax: PORTSCAN [-port=<port list>] [-timeout=<period in seconds>] [-host=<host name or IP>]")
}

func getCommandLineArguments () int {
	var ports string

    // host
    flag.StringVar (&host, "host", defaultHostIP, "Specifies host URL or IP.")

    //timeout
	portTimeout := flag.Int ("timeout", defaultPortTimeout, "Time allowed for TCP response (in seconds).")

   	if (*portTimeout <= 0) {
		*portTimeout = 1
	}

    fmt.Printf ("\nHost:'%s' (Timeout %ds).\n\n", host, *portTimeout)

    //port
	flag.StringVar(&ports, "port", defaultPort, "Specifies the ports to be scanned.")

    var allPorts []string

    fmt.Println(ports)

    if strings.Contains(ports,",") {
        allPorts = strings.Split(ports, ",")
    } else {
        allPorts = append(allPorts, ports)
    }

    fmt.Println(allPorts)

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
   
    fmt.Println(portList)

    return len(portList)
}

func tcpConnect(host string, port string) bool {
    conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), time.Duration(portTimeout))
    if err != nil {
        return false
    }
    if conn != nil {
        defer conn.Close()
       
        return true
    }

    return false
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
    countOpenPorts:=0
    for port,_ := range portList {
        portStr := fmt.Sprintf("%d", port)
        if tcpConnect (host, portStr) {
            fmt.Println("Opened", net.JoinHostPort(host, portStr))
            countOpenPorts++
        }
    }

    if countOpenPorts == 0 {
        fmt.Println("No open ports found.")
    }
}