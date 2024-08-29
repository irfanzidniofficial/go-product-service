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

func (r *shopRepository) GetShop(ctx context.Context, req *entity.GetShopRequest) (*entity.GetShopResponse, error) {
	var resp = new(entity.GetShopResponse)

	query := `
        SELECT name, description, terms
        FROM shops
        WHERE id =?
    `

	err := r.db.QueryRowxContext(ctx, r.db.Rebind(query), req.Id).StructScan(resp)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("respository::GetShop - Failed to get shop")
		return nil, err
	}

	return resp, nil
}

func (r *shopRepository) DeleteShop(ctx context.Context, req *entity.DeleteShopRequest) error {

	query := `
        UPDATE shops
        SET deleted_at = NOW()
        WHERE id = ? AND user_id= ?
    `

	_, err := r.db.ExecContext(ctx, r.db.Rebind(query), req.Id, req.UserId)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("respository::DeleteShop - Failed to delete shop")
		return err
	}

	return nil
}
