package context

import "github.com/gofiber/fiber/v2"

type Ctx struct {
	*fiber.Ctx
}

func NewCtx(ctx *fiber.Ctx) *Ctx {
	return &Ctx{
		Ctx: ctx,
	}
}

func (ctx *Ctx) SessionId() string {
	return ctx.Locals("SessionId").(string)
}

func (ctx *Ctx) IsLogon() bool {
	return ctx.Cookies("sid") != ""
}
