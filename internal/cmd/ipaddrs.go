package cmd

// ipAddrs contains ip addresses
type ipAddrs struct {
	m map[string]bool
}

// add adds an ip address to ipAddrs
func (a *ipAddrs) add(ip string) {
	a.m[ip] = true
}

// contains checks if ip is in ipAddrs
func (a *ipAddrs) contains(ip string) bool {
	return a.m[ip]
}

// newIPAddrs() creates a new ipAddrs
func newIPAddrs() *ipAddrs {
	return &ipAddrs{
		m: make(map[string]bool),
	}
}
