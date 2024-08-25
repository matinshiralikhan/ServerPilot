package agent

import (
	"log"
	"net"
	"time"
)

// Agent struct
type Agent struct {
	ServerAddress string
	Protocol      string
}

// NewAgent returns a new Agent instance
func NewAgent(serverAddress, protocol string) *Agent {
	return &Agent{
		ServerAddress: serverAddress,
		Protocol:      protocol,
	}
}

// StartAgent starts the agent to connect to the server
func (a *Agent) StartAgent() {
	for {
		conn, err := net.Dial(a.Protocol, a.ServerAddress)
		if err != nil {
			log.Printf("Failed to connect to server %s: %v\n", a.ServerAddress, err)
			time.Sleep(5 * time.Second)
			continue
		}

		log.Printf("Connected to server %s using %s protocol\n", a.ServerAddress, a.Protocol)

		// Add your communication logic here

		conn.Close()
	}
}
