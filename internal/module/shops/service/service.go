package service

import (
	"context"

	"go-product-service/internal/module/shops/entity"
	"go-product-service/internal/module/shops/ports"
)

var _ ports.ShopService = &shopService{}

type shopService struct {
	repo ports.ShopRepository
}

func NewShopService(repo ports.ShopRepository) *shopService {
	return &shopService{
		repo: repo,
	}
}

func (s *shopService) CreateShop(ctx context.Context, req *entity.CreateShopRequest) (*entity.CreateShopResponse, error) {
	return s.repo.CreateShop(ctx, req)
}
