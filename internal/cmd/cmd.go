package cmd

var (
	mcasts = newAddrs()
)

// Run is the main entry point
func Run() {
	runSocketLoop()
}
