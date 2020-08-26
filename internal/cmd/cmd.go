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
		log.Println("Accepting multicast address:", a)
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
		log.Println("Accepting destination port: ", p)
	}
}

func parseDests(dest string) {
	for _, d := range strings.Split(dest, ",") {
		// check if destination ip is valid
		if d == "" {
			continue
		}
		dst := newDest(d)
		if dst == nil {
			log.Fatal("invalid destination IP: ", d)
		}

		// add destination ip to list of destinations
		dests = append(dests, dst)
		log.Println("Forwarding to destination:  ", d)
	}
}

// parseCommandLine parses the command line arguments
func parseCommandLine() {
	var addresses = "224.0.0.251"
	var ports = "5353"
	var dest = ""

	// set command line arguments
	flag.StringVar(&addresses, "a", addresses,
		"only forward packets with this comma-separated list "+
			"of\nmulticast destination `addresses`, e.g.,\n"+
			"224.0.0.1,224.0.0.2")
	flag.StringVar(&ports, "p", ports,
		"only forward packets with this comma-separated list "+
			"of\ndestination `ports`, e.g., 1024,32000")
	flag.StringVar(&dest, "d", dest, "forward multicast packets to "+
		"this comma-separated list of\n`addresses`, "+
		"e.g., \"192.168.1.1,192.168.1.2\"")
	flag.Parse()

	// parse accepted multicast addresses
	if addresses == "" {
		log.Fatal("no multicast addresses specified")
	}
	parseAddresses(addresses)

	// parse accepted ports
	if ports == "" {
		log.Fatal("no ports specified")
	}
	parsePorts(ports)

	// parse destination addresses
	if dest == "" {
		log.Fatal("no destination IP specified")
	}
	parseDests(dest)
}

// Run is the main entry point
func Run() {
	parseCommandLine()
	runSocketLoop()
}
