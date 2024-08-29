package ports

import (
	"context"
	"go-product-service/internal/module/shops/entity"
)

type ShopRepository interface {
	CreateShop(ctx context.Context, req *entity.CreateShopRequest) (*entity.CreateShopResponse, error)
}

type ShopService interface {
	CreateShop(ctx context.Context, req *entity.CreateShopRequest) (*entity.CreateShopResponse, error)
}
