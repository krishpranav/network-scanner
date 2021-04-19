package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/krishpranav/network-scanner"
)

func PrintBanner() {

	fmt.Println("\n +----------------+")
	fmt.Println(" | Go PortScanner |")
	fmt.Println(" +----------------+")
}

// main function for scanning the port
func main() {

	var portList []string

	start := time.Noe()

	ports := flag.String("p", "80", "Port or ports to scan")
	all := flag.Bool("A", false, "Scans from port 1 to 1024")
	flag.Parse()

	if *all {
		*ports = "1-1024"
	}

	portsList = append(portsList, *ports)

	PrintBanner()

	elapsed := time.Since(start)
	fmt.Println("\nScanned in", elapsed)

}