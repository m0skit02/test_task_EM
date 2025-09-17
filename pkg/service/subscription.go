package service

import (
	"github.com/google/uuid"
	"time"
	"wb-task-L0/pkg/models"
	"wb-task-L0/pkg/repository"
)

type SubscriptionService struct {
	repo repository.Subscription
}

func NewSubscriptionService(repo repository.Subscription) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

func (s *SubscriptionService) Create(subscription *models.Subscription) (*models.Subscription, error) {
	uid, err := s.repo.Create(subscription)
	if err != nil {
		return nil, err
	}
	subscription.ID = uid
	return subscription, nil
}

func (s *SubscriptionService) GetAll() ([]models.Subscription, error) {
	return s.repo.GetAll()
}

func (s *SubscriptionService) GetByID(id string) (*models.Subscription, error) {
	return s.repo.GetByID(id)
}

func (s *SubscriptionService) Delete(id string) error {
	return s.repo.Delete(id)
}

// parseDateRange парсит даты формата "DD.MM.YYYY" и возвращает начало и конец дня
func parseDateRange(startStr, endStr string) (time.Time, time.Time, error) {
	const layout = "02.01.2006" // формат "день.месяц.год"

	start, err := time.Parse(layout, startStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	end, err := time.Parse(layout, endStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	// Начало дня и конец дня в UTC
	start = start.Truncate(24 * time.Hour)
	end = end.AddDate(0, 0, 1).Add(-time.Nanosecond) // до конца дня

	return start, end, nil
}

func (s *SubscriptionService) GetTotalCost(userIDStr, serviceName, startStr, endStr string) (int, error) {
	start, end, err := parseDateRange(startStr, endStr)
	if err != nil {
		return 0, err
	}

	var uid *uuid.UUID
	if userIDStr != "" {
		u, err := uuid.Parse(userIDStr)
		if err != nil {
			return 0, err
		}
		uid = &u
	}

	var svc *string
	if serviceName != "" {
		svc = &serviceName
	}

	return s.repo.GetTotalCost(uid, svc, start, end)
}
