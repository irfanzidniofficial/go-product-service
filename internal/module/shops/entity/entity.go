package entity

type CreateShopRequest struct {
	UserId      string `validate:"uuid" db:"user_id"`
	Name        string `validate:"required" db:"name"`
	Description string `validate:"required" db:"description"`
	Terms       string `validate:"required" db:"terms"`
}

type CreateShopResponse struct {
	Id string `json:"id" db:"id"`
}

type CreateShopResult struct {
}
