package healthcheck

type HealthcheckStatus int32

const (
	Unknown   HealthcheckStatus = 0
	Healthy   HealthcheckStatus = 1
	Unhealthy HealthcheckStatus = 2
)

func FromStatusCode(statusCode int) HealthcheckStatus {
	switch {
	case statusCode >= 200 && statusCode < 300:
		return Healthy
	case statusCode >= 400 && statusCode < 600:
		return Unhealthy
	default:
		return Unknown
	}
}

func (s HealthcheckStatus) String() string {
	switch s {
	case Healthy:
		return "Healthy"
	case Unhealthy:
		return "Unhealthy"
	default:
		return "Unknown"
	}
}
