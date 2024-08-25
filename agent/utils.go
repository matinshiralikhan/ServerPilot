package agent

import (
	"net"
	"time"
)

// PingServer checks the connectivity to a server
func PingServer(address string, timeout time.Duration) bool {
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
