package api

type Gauge struct {
	Name string `json:"name"`
	N    int64  `json:"n"`
}
