package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
	"wb-task-L0/pkg/models"
)

type Repository struct {
	Subscription
}

type Subscription interface {
	Create(subscription *models.Subscription) (uuid.UUID, error)
	GetAll() ([]models.Subscription, error)
	GetByID(id string) (*models.Subscription, error)
	Delete(id string) error
	GetTotalCost(userID *uuid.UUID, serviceName *string, startDate, endDate time.Time) (int, error)
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Subscription: NewSubscriptionRepo(db),
	}
}
