package runner

import (
	"net/http"
	"soupdevsolutions/healthchecker/domain"
)

func CheckHttpTarget(target domain.HealthcheckTarget) (domain.Healthcheck, error) {
	resp, err := http.Get(target.Uri)
	if err != nil {
		return domain.Healthcheck{}, err
	}

	return domain.Healthcheck{
		StatusCode: resp.StatusCode,
		Timestamp:  0,
	}, nil
}
