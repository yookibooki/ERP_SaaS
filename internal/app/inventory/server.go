package inventory

import (
	"log"
	"net/http"

	"github.com/yookibooki/ERP_SaaS/internal/pkg/config"
)

// Server holds the dependencies for our HTTP server.
type Server struct {
	config *config.Config
	// We'll add more dependencies like a database connection here later.
}

// NewServer creates a new instance of our server.
func NewServer(cfg *config.Config) *Server {
	return &Server{config: cfg}
}

// Run starts the HTTP server.
func (s *Server) Run() error {
	addr := s.config.Server.Addr
	log.Printf("Listening on http://localhost%s\n", addr)
	return http.ListenAndServe(addr, s.routes()) // Use our new router
}
