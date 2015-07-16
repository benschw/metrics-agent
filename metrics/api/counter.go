package api

type Counter struct {
	Name string `json:"name"`
	N    uint64 `json:"n"`
}
