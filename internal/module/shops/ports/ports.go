package ports

import (
	"context"
	"go-product-service/internal/module/shops/entity"
)

type ShopRepository interface {
	CreateShop(ctx context.Context, req *entity.CreateShopRequest) (*entity.CreateShopResponse, error)
	GetShop(ctx context.Context, req *entity.GetShopRequest) (*entity.GetShopResponse, error)
	DeleteShop(ctx context.Context, req *entity.DeleteShopRequest) error
}

type ShopService interface {
	CreateShop(ctx context.Context, req *entity.CreateShopRequest) (*entity.CreateShopResponse, error)
	GetShop(ctx context.Context, req *entity.GetShopRequest) (*entity.GetShopResponse, error)
	DeleteShop(ctx context.Context, req *entity.DeleteShopRequest) error
}
