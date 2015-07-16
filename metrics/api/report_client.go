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

func (c *ReportClient) FindAll() (Report, error) {
	var rep Report

	r, err := rest.MakeRequest("GET", fmt.Sprintf("http://%s/report", c.Address), nil)
	if err != nil {
		return rep, err
	}
	err = rest.ProcessResponseEntity(r, &rep, http.StatusOK)
	return rep, err
}

func (c *ReportClient) CounterAdd(name string, n uint64) error {
	ctr := &Counter{Name: name, N: n}

	r, err := rest.MakeRequest("POST", fmt.Sprintf("http://%s/counter", c.Address), ctr)
	if err != nil {
		return err
	}

	err = rest.ProcessResponseEntity(r, nil, http.StatusOK)
	return err
}
func (c *ReportClient) GaugeSet(name string, n int64) error {
	g := &Gauge{Name: name, N: n}

	r, err := rest.MakeRequest("POST", fmt.Sprintf("http://%s/gauge", c.Address), g)
	if err != nil {
		return err
	}

	err = rest.ProcessResponseEntity(r, nil, http.StatusOK)
	return err
}
