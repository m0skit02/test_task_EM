package service

import (
	"wb-task-EM/pkg/models"
	"wb-task-EM/pkg/repository"
)

type Subscription interface {
	Create(subscription *models.Subscription) (*models.Subscription, error)
	GetByID(id string) (*models.Subscription, error) // ← исправлено
	GetAll() ([]models.Subscription, error)
	Delete(id string) error
	GetTotalCost(userIDStr, serviceName, startStr, endStr string) (int, error)
}

type Service struct {
	Subscription
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Subscription: NewSubscriptionService(repos.Subscription),
	}
}
