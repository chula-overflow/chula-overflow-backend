package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleWare(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("sid")
	if cookie == "" {
		return ctx.SendStatus(fiber.ErrUnauthorized.Code)
	}
	ctx.Locals("SessionId", cookie)
	return ctx.Next()
}
