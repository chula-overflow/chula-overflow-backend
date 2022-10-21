package auth

import (
	"log"

	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/context"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/dto"
	"github.com/gofiber/fiber/v2"
)

type IService interface {
	Login(string) (*string, error)
	Revoke(string) error
}

type Handler struct {
	Service IService
}

// @Summary Login
// @Description Login with given email
// @Tags Auth
// @Param login body dto.Login false "email use for login"  Format(email)
// @Success 200
// @Router /auth/login [post]
func (h *Handler) Login(ctx *context.Ctx) {
	login := new(dto.Login)
	if err := ctx.BodyParser(login); err != nil {
		log.Print(err)
		ctx.JSON(err)
		return
	}

	res, err := h.Service.Login(login.Email)
	if err != nil {
		ctx.SendStatus(fiber.ErrServiceUnavailable.Code)
		log.Print(err)
		return
	}

	ctx.Cookie(&fiber.Cookie{
		Name:  "sid",
		Value: *res,
	})
}

// @Summary Revoke
// @Description Revoke session
// @Tags Auth
// @Success 200
// @Router /auth/revoke [get]
func (h *Handler) Revoke(ctx *context.Ctx) {
	err := h.Service.Revoke(ctx.SessionId())
	if err != nil {
		ctx.SendStatus(fiber.ErrServiceUnavailable.Code)
		log.Print(err)
		return
	}

	ctx.ClearCookie("sid")
}

func NewHandler(service IService) Handler {
	return Handler{
		Service: service,
	}
}
