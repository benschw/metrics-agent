package metrics

import (
	"log"
	"net/http"

	"github.com/benschw/metrics-agent/metrics/api"
	"github.com/benschw/opin-go/rest"
	"github.com/codahale/metrics"
)

type ReportResource struct {
}

func (r *ReportResource) GetAll(res http.ResponseWriter, req *http.Request) {
	counters, _ := metrics.Snapshot()

	if err := rest.SetOKResponse(res, counters); err != nil {
		rest.SetInternalServerErrorResponse(res, err)
		return
	}
}

func (r *ReportResource) Counter(res http.ResponseWriter, req *http.Request) {
	var ctr api.Counter

	if err := rest.Bind(req, &ctr); err != nil {
		log.Print(err)
		rest.SetBadRequestResponse(res)
		return
	}

	metrics.Counter(ctr.Name).AddN(ctr.N)

	if err := rest.SetOKResponse(res, ctr); err != nil {
		rest.SetInternalServerErrorResponse(res, err)
		return
	}
}
