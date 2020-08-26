package cmd

import (
	"encoding/binary"
	"log"
	"net"

	"golang.org/x/net/ipv4"
)

// runSocketLoop runs the main socket loop, reading packets from the socket
// and forwarding them to destination ip addresses
func runSocketLoop() {
	// open raw socket
	conn, err := net.ListenPacket("ip4:udp", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	raw, err := ipv4.NewRawConn(conn)
	if err != nil {
		log.Fatal(err)
	}

	// create packet buffer and start reading packets from raw socket
	log.Printf("Waiting for packets")
	buf := make([]byte, 2048)
	for {
		header, payload, _, err := raw.ReadFrom(buf)
		if err != nil {
			log.Fatal(err)
		}

		// print packet info
		destPort := binary.BigEndian.Uint16(payload[2:4])
		srcPort := binary.BigEndian.Uint16(payload[0:2])
		log.Printf("Got packet: %s:%d -> %s:%d\n", header.Src,
			srcPort, header.Dst, destPort)
	}
}
