package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
	"wb-task-L0/pkg/models"
)

type SubscriptionRepo struct {
	db *gorm.DB
}

func NewSubscriptionRepo(db *gorm.DB) *SubscriptionRepo {
	return &SubscriptionRepo{db: db}
}

func (r *SubscriptionRepo) Create(subscription *models.Subscription) (string, error) {
	if err := r.db.Create(subscription).Error; err != nil {
		return "", err
	}
	return subscription.ID.String(), nil
}

func (r *SubscriptionRepo) GetByID(id string) (*models.Subscription, error) {
	var subscription models.Subscription
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	if err := r.db.First(&subscription, "id = ?", uid).Error; err != nil {
		return nil, err
	}
	return &subscription, nil
}

func (r *SubscriptionRepo) GetAll() ([]models.Subscription, error) {
	var subscriptions []models.Subscription
	if err := r.db.Find(&subscriptions).Error; err != nil {
		return nil, err
	}
	return subscriptions, nil
}

func (r *SubscriptionRepo) Delete(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	if err := r.db.Delete(&models.Subscription{}, "id = ?", uid).Error; err != nil {
		return err
	}
	return nil
}

func (r *SubscriptionRepo) GetTotalCost(userID *uuid.UUID, serviceName *string, startDate, endDate time.Time) (int, error) {
	var total int64
	q := r.db.Model(&models.Subscription{}).
		Select("COALESCE(SUM(price), 0)")

	if userID != nil {
		q = q.Where("user_id = ?", *userID)
	}
	if serviceName != nil {
		q = q.Where("service_name = ?", *serviceName)
	}
	q = q.Where("start_date <= ? AND (end_date IS NULL OR end_date >= ?)", endDate, startDate)

	if err := q.Scan(&total).Error; err != nil {
		return 0, err
	}
	return int(total), nil
}
