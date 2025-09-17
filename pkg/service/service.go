package service

import "wb-task-L0/pkg/repository"

type Subscription interface {
	Create() error
}

type Service struct {
	Subscription
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Subscription: NewSubscriptionService(repos.Subscription),
	}
}
