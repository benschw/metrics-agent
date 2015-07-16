package api

type Report struct {
	Counters map[string]uint64 `json:"counters"`
	Gauges   map[string]int64  `json:"gauges"`
}
