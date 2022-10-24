package auth

import (
	"log"
	"time"

	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/context"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/dto"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IService interface {
	Login(*dto.Login) (*string, error)
	Revoke(string) error
	Me(string) (*dto.MeResponse, error)
}

type Handler struct {
	Service IService
}

// @Summary Login
// @Description Login with given email session are store in cookie as 'sid'. If no email were found in database, it will create one.
// @Description Cookies are automatically store in swag
// @Description return 200 if client already has sid on.
// @Tags Auth
// @Accept json
// @Param login body dto.Login false "email use for login"  Format(email)
// @Success 200
// @Failure 401
// @Failure 422
// @Router /auth/login [post]
func (h *Handler) Login(ctx *context.Ctx) error {
	if ctx.IsLogon() {
		return nil
	}

	login := new(dto.Login)
	if err := ctx.BodyParser(login); err != nil {
		// unprocessable entity
		return err
	}

	res, err := h.Service.Login(login)
	if err != nil {
		if _, ok := status.FromError(err); ok {
			log.Println("Service down")
			return fiber.ErrServiceUnavailable
		} else {
			log.Println("Parse error failed", err)
			return fiber.ErrInternalServerError
		}
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "sid",
		Value:   *res,
		Expires: time.Now().Add(24 * time.Hour),
	})

	return nil
}

// @Summary Revoke
// @Description Revoke session (expire session cookie from client)
// @Tags Auth
// @Success 200
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /auth/revoke [get]
func (h *Handler) Revoke(ctx *context.Ctx) error {
	err := h.Service.Revoke(ctx.SessionId())

	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				return fiber.ErrNotFound
			default:
				log.Println("Service down")
				return fiber.ErrServiceUnavailable
			}
		} else {
			log.Println("Parse error failed", err)
			return fiber.ErrInternalServerError
		}
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "sid",
		Expires:  time.Now().Add(-time.Second),
		Secure:   true,
		HTTPOnly: true,
	})

	return nil
}

// @Summary Me
// @Description Get current session detail
// @Tags Auth
// @Success 200 {object} dto.MeResponse
// @Failure 401
// @Failure 503
// @Failure 500
// @Router /auth/me [get]
func (h *Handler) Me(ctx *context.Ctx) error {
	ret, err := h.Service.Me(ctx.SessionId())
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				ctx.Cookie(&fiber.Cookie{
					Name:    "sid",
					Expires: time.Now().Add(-time.Second),
				})
				return nil
			default:
				log.Println("Service down")
				return fiber.ErrServiceUnavailable
			}
		} else {
			log.Println("Parse error failed", err)
			return fiber.ErrInternalServerError
		}
	}

	ctx.JSON(ret)
	return nil
}

func NewHandler(service IService) Handler {
	return Handler{
		Service: service,
	}
}
