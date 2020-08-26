package cmd

var (
	mcasts = newAddrs()
	dports = newAddrs()
	dests  []*dest
)

// Run is the main entry point
func Run() {
	runSocketLoop()
}
