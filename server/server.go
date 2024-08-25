package server

import (
	"log"
	"net"
	"net/http"
)

// Server struct
type Server struct {
	Address string
}

// NewServer returns a new Server instance
func NewServer(address string) *Server {
	return &Server{
		Address: address,
	}
}

// Start starts the server
func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.Address)
	if err != nil {
		log.Fatalf("Failed to start server on %s: %v", s.Address, err)
	}

	http.HandleFunc("/status", s.statusHandler)

	log.Printf("Server started on %s\n", s.Address)
	log.Fatal(http.Serve(listener, nil))
}

// statusHandler returns the server status
func (s *Server) statusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running"))
}
