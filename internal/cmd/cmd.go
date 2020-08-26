package cmd

var (
	// mcasts is a list of accepted multicast addresses
	mcasts = newAddrs()

	// dports is a list of accepted destination ports
	dports = newAddrs()

	// dests is the list of IPs to forward accepted packets to
	dests []*dest
)

// Run is the main entry point
func Run() {
	runSocketLoop()
}
