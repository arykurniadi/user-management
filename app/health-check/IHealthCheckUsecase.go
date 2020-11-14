package healthcheck

import "kriya.com/user-management/models"

type IHealthCheckUsecase interface {
	GetDBTimestamp() models.HealthCheck
}
