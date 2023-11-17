package viewmodels

import (
	"soupdevsolutions/healthchecker/healthcheck"
	"time"
)

type HealthcheckViewModel struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

type TargetViewModel struct {
	Id           string                 `json:"id"`
	Uri          string                 `json:"uri"`
	Name         string                 `json:"name"`
	LastStatus   string                 `json:"lastStatus"`
	Healthchecks []HealthcheckViewModel `json:"healthchecks"`
}

func NewTargetViewModel(target *healthcheck.HealthcheckTarget) TargetViewModel {
	var healthchecks []HealthcheckViewModel

	for _, hc := range target.Healthchecks {
		timestamp := time.Unix(hc.Timestamp, 0)

		healthchecks = append(healthchecks, HealthcheckViewModel{
			Status:    hc.Status.String(),
			Timestamp: timestamp.Format("2006-01-02 15:04:05"),
		})
	}

	return TargetViewModel{
		Id:           target.Id,
		Uri:          target.Uri,
		Name:         target.Name,
		LastStatus:   target.Status().String(),
		Healthchecks: healthchecks,
	}
}
