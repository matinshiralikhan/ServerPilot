package server

import (
	"log"
	"math/rand"
	"time"
)

// LoadBalancer struct
type LoadBalancer struct {
	Servers []string
}

// NewLoadBalancer returns a new LoadBalancer instance
func NewLoadBalancer(servers []string) *LoadBalancer {
	return &LoadBalancer{
		Servers: servers,
	}
}

// SelectServer selects the best server based on some criteria
func (lb *LoadBalancer) SelectServer() string {
	rand.Seed(time.Now().UnixNano())
	server := lb.Servers[rand.Intn(len(lb.Servers))]
	log.Printf("Selected server: %s\n", server)
	return server
}
