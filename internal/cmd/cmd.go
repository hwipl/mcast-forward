package cmd

var (
	mcasts = newAddrs()
	dports = newAddrs()
)

// Run is the main entry point
func Run() {
	runSocketLoop()
}
