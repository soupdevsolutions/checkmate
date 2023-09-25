package domain

type Healthcheck struct {
	Uri    string `json:"uri"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
