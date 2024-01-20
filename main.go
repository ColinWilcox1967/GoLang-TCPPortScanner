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
    defaultPort = ""
)

var (
    portTimeout int
    allPorts []string
    host string
)

func showTitle () {
	fmt.Printf ("\nTCP/IP Port Scanner (version %s)\n", portScannerVersion)
    fmt.Println("(c) 2021 Colin Wilcox")
}

func showSyntax () {
    fmt.Println ("Syntax: PORTSCAN [-port=<port list>] [-timeout=<period in seconds>] [-host=<host name or IP>]")
}

func getCommandLineArguments() int {
	var ports string

    // host
    flag.StringVar (&host, "host", defaultHostIP, "Specifies host URL or IP.")

    //timeout
	flag.IntVar(&portTimeout, "timeout", defaultPortTimeout, "Time allowed for TCP response (in seconds).")

    //port
	flag.StringVar(&ports, "port", defaultPort, "Specifies the ports to be scanned.")
  
     
    flag.Parse()

    if (portTimeout <= 0) {
		portTimeout = defaultPortTimeout
	}

    fmt.Printf ("\nHost:'%s' (Timeout %ds).\n\n", host, portTimeout)
    
    if ports == "" {
        // scan all ports 0 .. 65535
        for portNumber := 0; portNumber <= 65535; portNumber++ {
            allPorts = append(allPorts, fmt.Sprintf("%d", portNumber))
        }
    } else {
        // parse port argument A,B,C or A-B
        if strings.Contains(ports,",") {
            // -ports=A,B, ... ,D
            allPorts = strings.Split(ports, ",")
           
        } else 
        if strings.Contains(ports, "-") {

            // -ports=A-B
            portRange := strings.Split(ports, "-")
            firstPort,_ := strconv.Atoi(portRange[0])
            lastPort,_ := strconv.Atoi(portRange[1])

            for portNumber := firstPort; portNumber <= lastPort; portNumber++ {
                allPorts = append(allPorts, fmt.Sprintf("%d", portNumber))
            }
        
        } else {
            allPorts = append(allPorts, ports)
        }
    }

    return len(allPorts)
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

    fmt.Println ("Scanning ports ...")
    for _, portStr := range allPorts {
      //  port := fmt.Sprintf("%d", portStr)
        if tcpConnect (host, portStr) {
            fmt.Println("Open port : ", net.JoinHostPort(host, portStr))
            countOpenPorts++
        }
    }
   
    if countOpenPorts == 0 {
        fmt.Println("No open ports found.")
    }
}