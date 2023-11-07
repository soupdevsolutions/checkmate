package healthcheck

type HealthcheckTarget struct {
	Id           string
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
	Id        string
	Status    HealthcheckStatus `json:"status"`
	Timestamp int64             `json:"timestamp"`
}
