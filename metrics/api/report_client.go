package api

import (
	"fmt"
	"net/http"

	"github.com/benschw/opin-go/rest"
)

// Client Factory
func NewReportClient() *ReportClient {
	return &ReportClient{
		Address: "localhost:8080",
	}
}

// Client
type ReportClient struct {
	Address string
}

func (c *ReportClient) FindAll() (map[string]uint64, error) {
	var ctr map[string]uint64

	r, err := rest.MakeRequest("GET", fmt.Sprintf("http://%s/report", c.Address), nil)
	if err != nil {
		return ctr, err
	}
	err = rest.ProcessResponseEntity(r, &ctr, http.StatusOK)
	return ctr, err
}

func (c *ReportClient) CounterAdd(name string, n uint64) error {
	ctr := &Counter{Name: name, N: n}

	r, err := rest.MakeRequest("GET", fmt.Sprintf("http://%s/counter", c.Address), ctr)
	if err != nil {
		return err
	}

	err = rest.ProcessResponseEntity(r, nil, http.StatusOK)
	return err
}
