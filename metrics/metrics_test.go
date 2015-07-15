package metrics

import (
	"github.com/benschw/metrics-agent/metrics/api"
	"github.com/codahale/metrics"
	. "gopkg.in/check.v1"
)

var _ = Suite(&TestSuite{})

type TestSuite struct {
	svc    *MetricsService
	client *api.ReportClient
}

func (s *TestSuite) SetUpSuite(c *C) {
	s.client, s.svc = GetClientAndService()

	go s.svc.Run()
}
func (s *TestSuite) TearDownSuite(c *C) {
	s.svc.Stop()
}
func (s *TestSuite) SetUpTest(c *C) {
	metrics.Reset()
}

func (s *TestSuite) TestCounter(c *C) {
	// given
	var expected = make(map[string]uint64, 0)

	// when
	s.client.CounterAdd("foo", 1)
	found, err := s.client.FindAll()

	// then
	c.Assert(err, Equals, nil)

	c.Assert(found, DeepEquals, expected)
}
