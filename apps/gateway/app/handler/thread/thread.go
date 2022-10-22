package auth

import (
	"log"

	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/context"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/dto"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IService interface {
	GetThread(string) (*dto.Thread, error)
	CreateThread(*dto.CreateThread) (*dto.CreateThreadResponse, error)
	CreateReply(*dto.CreateReply, string) error
}

type Handler struct {
	Service IService
}

// @Summary GetThread
// @Description get specific thread
// @Tags Thread
// @Produce json
// @Accept json
// @Param thread_id path string true "Thread id"
// @Success 200 {object} dto.Thread
// @Failure 500
// @Failure 503
// @Router /thread/:thread_id [get]
func (h *Handler) GetThread(ctx *context.Ctx) error {
	threadId := ctx.Params("thread_id")

	thread, err := h.Service.GetThread(threadId)
	if err != nil {
		if _, ok := status.FromError(err); ok {
			log.Println("Service down")
			return fiber.ErrServiceUnavailable
		} else {
			log.Println("Parse error failed", err)
			return fiber.ErrInternalServerError
		}
	}

	log.Print(thread)
	ctx.JSON(thread)

	return nil
}

// @Summary CreateThread
// @Description Create thread (require session)
// @Tags Thread
// @Accept json
// @Param thread body dto.CreateThread true "Thread information"
// @Success 200 {object} dto.CreateThreadResponse
// @Failure 401
// @Failure 500
// @Failure 503
// @Router /thread [post]
func (h *Handler) CreateThread(ctx *context.Ctx) error {
	createThread := new(dto.CreateThread)

	if err := ctx.BodyParser(&createThread); err != nil {
		return err
	}

	// todo add user id
	response, err := h.Service.CreateThread(createThread)

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

	ctx.JSON(response)

	return nil
}

// @Summary CreateReply
// @Description Create reply in thread (require session)
// @Tags Thread
// @Accept json
// @Param reply body dto.CreateReply true "Reply information"
// @Param thread_id path string true "Thread id"
// @Success 200
// @Failure 401
// @Failure 500
// @Failure 503
// @Router /thread/:thread_id/reply [post]
func (h *Handler) CreateReply(ctx *context.Ctx) error {
	createReply := new(dto.CreateReply)

	if err := ctx.BodyParser(&createReply); err != nil {
		return err
	}

	threadId := ctx.Params("thread_id")

	// todo add user id
	err := h.Service.CreateReply(createReply, threadId)

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

	return nil
}

func NewHandler(service IService) Handler {
	return Handler{
		Service: service,
	}
}
