package middleware

import (
	"log"

	authSrv "github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/service/auth"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/validator"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateMiddleWare(srv *authSrv.Service, validator *validator.MyValidator) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		cookie := ctx.Cookies("sid")
		if cookie == "" {
			return fiber.ErrUnauthorized
		}

		err := validator.Var(cookie, "uuid4")
		if err != nil {
			return fiber.ErrForbidden
		}

		userId, err := srv.Validate(cookie)
		if err != nil {
			if e, ok := status.FromError(err); ok {
				switch e.Code() {
				case codes.NotFound:
					return fiber.ErrForbidden
				default:
					log.Println("Service down")
					return fiber.ErrServiceUnavailable
				}
			} else {
				log.Println("Parse error failed", err)
				return fiber.ErrInternalServerError
			}
		}

		ctx.Locals("UserId", userId)
		ctx.Locals("SessionId", cookie)
		return ctx.Next()
	}
}
