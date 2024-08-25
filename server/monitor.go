package server

import (
	"log"
	"time"
	"github.com/username/ServerPilot/agent"
)

// Monitor struct
type Monitor struct {
	Servers []string
}

// NewMonitor returns a new Monitor instance
func NewMonitor(servers []string) *Monitor {
	return &Monitor{
		Servers: servers,
	}
}

// StartMonitoring starts monitoring the servers
func (m *Monitor) StartMonitoring() {
	for {
		for _, server := range m.Servers {
			if !agent.PingServer(server, 2*time.Second) {
				log.Printf("Server %s is down!\n", server)
			} else {
				log.Printf("Server %s is up and running\n", server)
			}
		}
		time.Sleep(10 * time.Second)
	}
}
