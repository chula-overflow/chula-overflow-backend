package auth

import (
	"log"

	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/validator"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/context"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/dto"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IService interface {
	GetExam(string) (*dto.Exam, error)
}

type Handler struct {
	Service IService
	v       *validator.MyValidator
}

// @Summary GetExam
// @Description Get specific exam
// @Tags Exam
// @Param exam_id path string true "exam ID"
// @Success 200 {object} dto.Exam
// @Failure 404
// @Failure 500
// @Failure 503
// @Router /exam/:exam_id [get]
func (h *Handler) GetExam(ctx *context.Ctx) error {
	examId := ctx.Params("exam_id")

	exam, err := h.Service.GetExam(examId)

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

	ctx.JSON(exam)

	return nil
}

func NewHandler(service IService, v *validator.MyValidator) Handler {
	return Handler{
		Service: service,
		v:       v,
	}
}
