package service

import "github.com/Icorp/webStore/pkg/repository"

// Business logics

type Authorization interface {
}

type ProductList interface {
}

type ProductItem interface {
}

type Service struct {
	Authorization
	ProductItem
	ProductList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
