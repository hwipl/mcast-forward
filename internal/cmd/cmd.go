package cmd

import (
	"flag"
	"log"
	"net"
	"strconv"
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

func parsePorts(ports string) {
	for _, p := range strings.Split(ports, ",") {
		// check if port is valid
		_, err := strconv.ParseUint(p, 10, 16)
		if err != nil {
			log.Fatal("error parsing port ", p, ": ", err)
		}

		// add port to accepted ports
		dports.add(p)
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
	flag.StringVar(&ports, "p", ports,
		"only forward packets with this comma-separated list "+
			"of\ndestination `ports`, e.g., 1024,32000")
	flag.Parse()

	// parse ip addresses
	if addresses != "" {
		parseAddresses(addresses)
	}

	// parse ports
	if ports != "" {
		parsePorts(ports)
	}
}

// Run is the main entry point
func Run() {
	parseCommandLine()
	runSocketLoop()
}
