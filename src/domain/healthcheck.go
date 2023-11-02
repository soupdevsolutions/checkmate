package domain

type HealthcheckTarget struct {
	Uri          string        `json:"uri"`
	Name         string        `json:"name"`
	Healthchecks []Healthcheck `json:"healthchecks"`
}

type Healthcheck struct {
	StatusCode int   `json:"statusCode"`
	Timestamp  int64 `json:"timestamp"`
}
