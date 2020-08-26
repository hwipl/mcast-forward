package cmd

import (
	"flag"
	"log"
	"net"
	"strings"
)

var (
	// mcasts is a list of accepted multicast addresses
	mcasts = newAddrs()

	// dports is a list of accepted destination ports
	dports = newAddrs()

	// dests is the list of IPs to forward accepted packets to
	dests []*dest
)

func parseAddresses(addresses string) {
	for _, a := range strings.Split(addresses, ",") {
		// check if ip address is valid
		ip := net.ParseIP(a)
		if ip == nil {
			log.Fatal("cannot parse address ", a)
		}
		if ip.To4() == nil {
			log.Fatal("address ", a, " is not an IPv4 address")
		}
		if !ip.IsMulticast() {
			log.Fatal("address ", a, " is not a multicast address")
		}

		// add address to accepted multicast addresses
		mcasts.add(a)
	}
}

// parseCommandLine parses the command line arguments
func parseCommandLine() {
	var addresses = "224.0.0.1"
	var ports = "6112"

	// set command line arguments
	flag.StringVar(&addresses, "a", addresses,
		"only forward packets with this comma-separated list "+
			"of\nmulticast `addresses`, e.g., 224.0.0.1,224.0.0.2")
	flag.Parse()

	// parse ip addresses
	if addresses != "" {
		parseAddresses(addresses)
	}
}

// Run is the main entry point
func Run() {
	parseCommandLine()
	runSocketLoop()
}
