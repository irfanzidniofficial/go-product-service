package handler

import (
	"go-product-service/internal/adapter"
	"go-product-service/internal/middleware"
	"go-product-service/internal/module/shops/entity"
	"go-product-service/internal/module/shops/ports"
	"go-product-service/internal/module/shops/repository"
	"go-product-service/internal/module/shops/service"
	"go-product-service/pkg/errmsg"
	"go-product-service/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type shopHandler struct {
	service ports.ShopService
}

func NewShopHandler() *shopHandler {
	var (
		handler = new(shopHandler)
		repo    = repository.NewShopRepository(adapter.Adapters.ShopeefunPostgres)
		service = service.NewShopService(repo)
	)
	handler.service = service
	return handler
}

func (h *shopHandler) Register(router fiber.Router) {
	router.Post("/shops", middleware.UserIdHeader, h.CreateShop)

}

func (h *shopHandler) CreateShop(c *fiber.Ctx) error {
	var (
		req = new(entity.CreateShopRequest)
		ctx = c.Context()
		v   = adapter.Adapters.Validator
		l   = middleware.GetLocals(c)
	)

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::CreateShop - Parse request body")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	req.UserId = l.UserId

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("payload", req).Msg("handler::CreateShop - Validate request body")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.CreateShop(ctx, req)
	if err != nil {
		log.Error().Err(err).Any("payload", req).Msg("handler::CreateShop - CreateShop")
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}
	return c.Status(fiber.StatusCreated).JSON(response.Success(resp, ""))

}
