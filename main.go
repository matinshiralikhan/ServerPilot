package main

import (
	"github.com/username/ServerPilot/agent"
	"github.com/username/ServerPilot/server"
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Servers  []string `json:"servers"`
	Protocol string   `json:"protocol"`
	LogFile  string   `json:"log_file"`
}

func main() {
	// Load config
	configData, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		log.Fatalf("Unable to read config file: %v", err)
	}

	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalf("Unable to parse config file: %v", err)
	}

	// Start Monitoring
	monitor := server.NewMonitor(config.Servers)
	go monitor.StartMonitoring()

	// Start Load Balancer
	lb := server.NewLoadBalancer(config.Servers)
	selectedServer := lb.SelectServer()

	// Start Agent
	agent := agent.NewAgent(selectedServer, config.Protocol)
	go agent.StartAgent()

	// Start Server
	srv := server.NewServer(":8080")
	srv.Start()
}
