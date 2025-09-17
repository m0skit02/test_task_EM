package repository

import (
	"gorm.io/gorm"
	"wb-task-L0/internal/models"
)

type Repository struct {
	Subscription
}

type Subscription interface {
	Create(subscription models.Subscription) (models.Subscription, error)
	GetAll() ([]models.Subscription, error)
	GetByID(id string) (models.Subscription, error)
	Delete(id string) error
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Subscription: NewSubscriptionRepo{db},
	}
}
