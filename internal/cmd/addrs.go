package cmd

// addrs is a set of addresses
type addrs struct {
	m map[string]bool
}

// add adds addr to addrs
func (a *addrs) add(addr string) {
	a.m[addr] = true
}

// contains checks if addr is in addrs
func (a *addrs) contains(addr string) bool {
	return a.m[addr]
}

// newAddrs() creates a new addrs
func newAddrs() *addrs {
	return &addrs{
		m: make(map[string]bool),
	}
}
