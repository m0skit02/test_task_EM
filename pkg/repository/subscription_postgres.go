package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
	"wb-task-EM/pkg/models"
)

type SubscriptionRepo struct {
	db *gorm.DB
}

func NewSubscriptionRepo(db *gorm.DB) *SubscriptionRepo {
	return &SubscriptionRepo{db: db}
}

func (r *SubscriptionRepo) Create(subscription *models.Subscription) (uuid.UUID, error) {
	if err := r.db.Create(subscription).Error; err != nil {
		return uuid.Nil, err
	}
	return subscription.ID, nil
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

	sql := `
	SELECT COALESCE(SUM(price * (
		DATE_PART('year', AGE(LEAST(COALESCE(end_date, $1), $2), GREATEST(start_date, $3))) * 12 +
		DATE_PART('month', AGE(LEAST(COALESCE(end_date, $1), $2), GREATEST(start_date, $3)))
	)), 0)::bigint AS total
	FROM subscriptions
	WHERE ($4::uuid IS NULL OR user_id = $4)
	  AND ($5::text IS NULL OR service_name = $5)
	  AND start_date <= $6
	  AND (end_date IS NULL OR end_date >= $7)
	`

	args := []interface{}{
		endDate,
		endDate,
		startDate,
		userID,
		serviceName,
		endDate,
		startDate,
	}

	if err := r.db.Raw(sql, args...).Scan(&total).Error; err != nil {
		return 0, err
	}

	return int(total), nil
}
