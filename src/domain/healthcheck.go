package domain

type HealthcheckTarget struct {
	Uri          string        `json:"uri"`
	Name         string        `json:"name"`
	Healthchecks []Healthcheck `json:"healthchecks"`
}

func NewHealthcheckTarget(name string, uri string) HealthcheckTarget {
	return HealthcheckTarget{
		Uri:          uri,
		Name:         name,
		Healthchecks: []Healthcheck{},
	}
}

type Healthcheck struct {
	StatusCode int   `json:"statusCode"`
	Timestamp  int64 `json:"timestamp"`
}
