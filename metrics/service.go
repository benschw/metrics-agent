package metrics

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/benschw/opin-go/ophttp"
	"github.com/gorilla/mux"
)

func NewMetricsService(bind string) (*MetricsService, error) {
	server := ophttp.NewServer(bind)

	return &MetricsService{Server: server}, nil
}

type MetricsService struct {
	Server *ophttp.Server
}

// Configure and start http server
func (s *MetricsService) Run() error {
	sock, err := net.Listen("tcp", "localhost:8123")
	if err != nil {
		return err
	}
	go func() {
		fmt.Println("HTTP now available at port 8123")
		http.Serve(sock, nil)
	}()

	// Build Resource
	report := &ReportResource{}

	// Wire Routes
	r := mux.NewRouter()
	r.HandleFunc("/report", report.GetAll).Methods("GET")
	r.HandleFunc("/counter", report.Counter).Methods("POST")

	http.Handle("/", r)

	// Start Server
	err = s.Server.Start()

	log.Println("Server Stopped")
	return err
}

func (s *MetricsService) Stop() {
	log.Println("Stopping Server...")
	s.Server.Stop()
}
