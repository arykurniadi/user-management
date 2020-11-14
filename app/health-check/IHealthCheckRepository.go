package healthcheck

import "kriya.com/user-management/models"

type IHealthCheckRepository interface {
	GetDBTimestamp() models.HealthCheck
}
