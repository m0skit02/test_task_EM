package repository

import (
	"gorm.io/gorm"
	"wb-task-L0/internal/models"
)

type SubscriptionRepo struct {
	db *gorm.DB
}

func NewSubscriptionRepo(db *gorm.DB) *SubscriptionRepo {
	return &SubscriptionRepo{db: db}
}

func (r *SubscriptionRepo) Create(subscription *models.Subscription) (string, error) {}

func (r *SubscriptionRepo) GetByID(userID string) (models.Subscription, error) {}

func (r *SubscriptionRepo) GetALL() ([]models.Subscription, error) {}

func (r *SubscriptionRepo) Delete(userID string) error {}
