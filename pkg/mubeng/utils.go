package mubeng

import "crypto/tls"

// getUnsafeCipherSuites returns a list of all cipher suites that are considered
// unsafe by the Go standard library. This includes all cipher suites that are
// not included in the default cipher suite list.
func getUnsafeCipherSuites() []uint16 {
	// Copied from: https://github.com/projectdiscovery/nuclei/pull/4753/files
	unsafeCipherSuites := make([]uint16, 0, len(tls.InsecureCipherSuites())+len(tls.CipherSuites()))
	for _, suite := range tls.InsecureCipherSuites() {
		unsafeCipherSuites = append(unsafeCipherSuites, suite.ID)
	}
	for _, suite := range tls.CipherSuites() {
		unsafeCipherSuites = append(unsafeCipherSuites, suite.ID)
	}

	return unsafeCipherSuites
}
