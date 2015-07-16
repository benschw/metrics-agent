package metrics

import (
	"log"

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
	var expected = api.Report{
		Counters: map[string]uint64{"foo": 3},
		Gauges:   map[string]int64{},
	}

	// when
	s.client.CounterAdd("foo", 1)
	s.client.CounterAdd("foo", 2)
	found, err := s.client.FindAll()

	// then
	c.Assert(err, Equals, nil)

	c.Assert(found, DeepEquals, expected)
}

func (s *TestSuite) TestGauge(c *C) {
	// given
	var expected = api.Report{
		Counters: map[string]uint64{},
		Gauges:   map[string]int64{"foo": 15},
	}

	// when
	s.client.GaugeSet("foo", 10)
	s.client.GaugeSet("foo", 15)
	found, err := s.client.FindAll()

	// then
	c.Assert(err, Equals, nil)

	c.Assert(found, DeepEquals, expected)
}

func (s *TestSuite) TestHistogram(c *C) {

	expected := map[string]int64{
		"foo.P50":  5,
		"foo.P75":  5,
		"foo.P90":  15,
		"foo.P95":  15,
		"foo.P99":  15,
		"foo.P999": 15,
	}

	h := metrics.NewHistogram("foo", 0, 100, 4)

	h.RecordValue(1)
	h.RecordValue(5)
	h.RecordValue(5)
	h.RecordValue(15)

	_, g := metrics.Snapshot()
	log.Printf("g: %+v", g)

	c.Assert(g, DeepEquals, expected)
}
