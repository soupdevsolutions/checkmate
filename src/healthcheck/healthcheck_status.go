package healthcheck

type HealthcheckStatus int32

const (
	UNKNOWN   HealthcheckStatus = 0
	HEALTHY   HealthcheckStatus = 1
	UNHEALTHY HealthcheckStatus = 2
)

func FromStatusCode(statusCode int) HealthcheckStatus {
	switch {
	case statusCode >= 200 && statusCode < 300:
		return HEALTHY
	case statusCode >= 400 && statusCode < 600:
		return UNHEALTHY
	default:
		return UNKNOWN
	}
}

func (s HealthcheckStatus) String() string {
	switch s {
	case HEALTHY:
		return "HEALTHY"
	case UNHEALTHY:
		return "UNHEALTHY"
	default:
		return "UNKNOWN"
	}
}
