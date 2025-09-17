package service

import (
	"github.com/google/uuid"
	"time"
	"wb-task-L0/pkg/repository"
)

type SubscriptionService struct {
	repo repository.Subscription
}

func NewSubscriptionService(repo repository.Subscription) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

func (s *SubscriptionService) Create() error {}

func (s *SubscriptionService) GetAll() ([]Subscription, error) {}

func (s *SubscriptionService) GetByID() {}

func (s *SubscriptionService) Delete(ID string) error {}

// pkg/service/subscription.go
func (s *SubscriptionService) GetTotalCost(userIDStr, serviceName, startStr, endStr string) (int, error) {
	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		return 0, err
	}
	end, err := time.Parse("2006-01-02", endStr)
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
