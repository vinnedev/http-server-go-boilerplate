/**
* This is a layer of business logic for health check
**/

package app

type HealthService struct{}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (s *HealthService) CheckHealth() map[string]string {
	return map[string]string{"status": "healthy"}
}
