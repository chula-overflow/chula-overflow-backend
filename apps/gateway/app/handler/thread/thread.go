package auth

import (
	"log"
	"strconv"

	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/validator"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/context"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/dto"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IService interface {
	DownvoteThread(threadId string) (*dto.ThreadBody, error)
	UpvoteThread(threadId string) (*dto.ThreadBody, error)
	GetThreadById(threadId string) (*dto.ThreadBody, error)
	GetAllThreadsByExamProperty(year int32, semester string, term string) ([]*dto.ThreadBody, error)
	CreateThread(userId string, body *dto.ThreadRequestCreateBody) (*dto.ThreadBody, error)
}

type Handler struct {
	Service IService
	v       *validator.MyValidator
}

// @Summary Get thread by id
// @Tags Thread
// @Produce json
// @Param thread_id path string true "Thread id"
// @Success 200 {object} dto.ThreadBody
// @Success 404
// @Router /thread/:thread_id [get]
func (h *Handler) GetThreadById(ctx *context.Ctx) error {
	threadId := ctx.Params("thread_id")

	thread, err := h.Service.GetThreadById(threadId)
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

	ctx.JSON(thread)

	return nil
}

// @Summary Get thread by property
// @Tags Thread
// @Accept json
// @Param year query int true "year"
// @Param semester query string true "semester"
// @Param term query string true "term"
// @Success 200 {object} []dto.ThreadBody
// @Failure 400
// @Router /thread [get]
func (h *Handler) GetAllThreadsByExamProperty(ctx *context.Ctx) error {
	year := ctx.Query("year")
	semester := ctx.Query("semester")
	term := ctx.Query("term")

	if year == "" || semester == "" || term == "" {
		return fiber.ErrBadRequest
	}

	yearInt, err := strconv.Atoi(year)

	if err != nil {
		return fiber.ErrBadRequest
	}

	res, err := h.Service.GetAllThreadsByExamProperty(int32(yearInt), semester, term)

	if err != nil {
		return err
	}

	ctx.JSON(res)

	return nil
}

// @Summary Create thread
// @Tags Thread
// @Produce json
// @Param thread_body body dto.ThreadRequestCreateBody true "Thread body"
// @Success 200 {object} dto.ThreadBody
// @Success 404
// @Router /thread [post]
func (h *Handler) CreateThread(ctx *context.Ctx) error {
	createThread := new(dto.ThreadRequestCreateBody)

	userId := ctx.UserId()

	if err := ctx.BodyParser(&createThread); err != nil {
		return err
	}

	if err := h.v.Struct(createThread); err != nil {
		return fiber.ErrBadRequest
	}

	thread, err := h.Service.CreateThread(userId, createThread)
	if err != nil {
		return err
	}

	ctx.JSON(thread)

	return nil
}

// @Summary Upvote thread
// @Tags Thread
// @Produce json
// @Param thread_id path string true "thread id"
// @Success 200 {object} dto.ThreadBody
// @Failure 400
// @Router /thread/:thread_id/upvote [post]
func (h *Handler) UpvoteThread(ctx *context.Ctx) error {
	threadId := ctx.Params("thread_id")

	res, err := h.Service.UpvoteThread(threadId)

	if err != nil {
		return err
	}

	ctx.JSON(res)

	return nil
}

// @Summary Downvote thread
// @Tags Thread
// @Produce json
// @Param thread_id path string true "thread id"
// @Success 200 {object} dto.ThreadBody
// @Failure 400
// @Router /thread/:thread_id/downvote [post]
func (h *Handler) DownvoteThread(ctx *context.Ctx) error {
	threadId := ctx.Params("thread_id")

	res, err := h.Service.DownvoteThread(threadId)

	if err != nil {
		return err
	}

	ctx.JSON(res)

	return nil
}
func NewHandler(service IService, v *validator.MyValidator) Handler {
	return Handler{
		Service: service,
		v:       v,
	}
}
