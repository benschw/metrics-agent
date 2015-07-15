package metrics

import (
	"fmt"
	"testing"

	"github.com/benschw/dns-clb-go/dns"
	"github.com/benschw/metrics-agent/metrics/api"
	"github.com/benschw/opin-go/ophttp"
	"github.com/benschw/opin-go/rando"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

func GetClientAndService() (*api.ReportClient, *MetricsService) {
	address := dns.Address{Address: "localhost", Port: uint16(rando.Port())}

	server := ophttp.NewServer(fmt.Sprintf("%s:%d", address.Address, address.Port))
	svc := &MetricsService{
		Server: server,
	}

	client := &api.ReportClient{
		Address: address.String(),
	}
	return client, svc
}
