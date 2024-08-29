package repository

import (
	"context"
	"go-product-service/internal/module/shops/entity"
	"go-product-service/internal/module/shops/ports"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

var _ ports.ShopRepository = &shopRepository{}

type shopRepository struct {
	db *sqlx.DB
}

func NewShopRepository(db *sqlx.DB) *shopRepository {
	return &shopRepository{
		db: db,
	}
}

func (r *shopRepository) CreateShop(ctx context.Context, req *entity.CreateShopRequest) (*entity.CreateShopResponse, error) {

	var resp = new(entity.CreateShopResponse)

	query := `
		INSERT INTO shops (user_id, name, description, terms)
		VALUES (?, ?, ?, ?) RETURNING id
	`

	err := r.db.QueryRowContext(ctx, r.db.Rebind(query),
		req.UserId,
		req.Name,
		req.Description,
		req.Terms).Scan(&resp.Id)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("respository::CreateShop - Failed to create shop")
		return nil, err
	}

	return resp, nil

}
