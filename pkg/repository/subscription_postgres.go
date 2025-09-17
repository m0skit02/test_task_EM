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

	// Подзапрос для вычисления пересекающегося периода
	subQuery := r.db.Model(&models.Subscription{}).
		Select(`
			price,
			GREATEST(start_date, ?) AS greatest_gstart,
			LEAST(COALESCE(end_date, ?), ?) AS least_lend
		`, startDate, endDate, endDate).
		Where("start_date <= ? AND (end_date IS NULL OR end_date >= ?)", endDate, startDate)

	// Фильтры
	if userID != nil {
		subQuery = subQuery.Where("user_id = ?", *userID)
	}
	if serviceName != nil {
		subQuery = subQuery.Where("service_name = ?", *serviceName)
	}

	// Основной запрос: суммируем цену с учетом количества месяцев пересечения
	sql := `
	SELECT COALESCE(SUM(price * (
		(DATE_PART('year', least_lend) - DATE_PART('year', greatest_gstart)) * 12
		+ (DATE_PART('month', least_lend) - DATE_PART('month', greatest_gstart)) + 1
	)), 0)::bigint AS total
	FROM (?) AS s
	WHERE greatest_gstart <= least_lend
	`

	if err := r.db.Raw(sql, subQuery).Scan(&total).Error; err != nil {
		return 0, err
	}

	return int(total), nil
}
