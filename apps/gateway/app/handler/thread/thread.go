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
	CreateThread(userId string, body *dto.ThreadRequestCreateBody) (*dto.ThreadBody, error)
	GetThread(year int32, semester string, term string, courseId string) ([]*dto.ThreadBody, error)
	GetThreadById(threadId string) (*dto.ThreadBody, error)

	DownvoteThread(threadId string) (*dto.ThreadBody, error)
	UpvoteThread(threadId string) (*dto.ThreadBody, error)

	DownvoteAnswer(answerId string) (*dto.AnswerBody, error)
	UpvoteAnswer(answerId string) (*dto.AnswerBody, error)

	DownvoteProblem(problemId string) (*dto.ProblemBody, error)
	UpvoteProblem(problemId string) (*dto.ProblemBody, error)

	AddAnswer(user string, threadId string, body *dto.AnswerRequestCreateBody) (*dto.AnswerBody, error)
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
// @Param year query int false "Year"
// @Param semester query string false "Semester"
// @Param term query string false "Term"
// @Param course_id query string false "Course Id"
// @Success 200 {object} []dto.ThreadBody
// @Router /thread [get]
func (h *Handler) GetThread(ctx *context.Ctx) error {
	year := ctx.Query("year")
	semester := ctx.Query("semester")
	term := ctx.Query("term")
	courseId := ctx.Query("course_id")

	yearInt, err := strconv.Atoi(year)

	if err != nil {
		return fiber.ErrBadRequest
	}

	res, err := h.Service.GetThread(int32(yearInt), semester, term, courseId)

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
// @Success 400
// @Success 401
// @Success 404
// @Success 422
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
// @Param thread_id path string true "Thread id"
// @Success 200 {object} dto.ThreadBody
// @Success 401
// @Success 404
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
// @Success 401
// @Success 404
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

// @Summary Upvote answer
// @Tags Thread
// @Produce json
// @Param answer_id path string true "Answer id"
// @Success 200 {object} dto.AnswerBody
// @Success 401
// @Success 404
// @Router /answer/:answer_id/upvote [post]
func (h *Handler) UpvoteAnswer(ctx *context.Ctx) error {
	answerId := ctx.Params("answer_id")

	res, err := h.Service.UpvoteAnswer(answerId)

	if err != nil {
		return err
	}

	ctx.JSON(res)

	return nil
}

// @Summary Downvote answer
// @Tags Thread
// @Produce json
// @Param answer_id path string true "Answer id"
// @Success 200 {object} dto.AnswerBody
// @Failure 400
// @Success 401
// @Success 404
// @Router /answer/:answer_id/downvote [post]
func (h *Handler) DownvoteAnswer(ctx *context.Ctx) error {
	answerId := ctx.Params("answer_id")

	res, err := h.Service.DownvoteAnswer(answerId)

	if err != nil {
		return err
	}

	ctx.JSON(res)

	return nil
}

// @Summary Upvote problem
// @Tags Thread
// @Produce json
// @Param problem_id path string true "Problem id"
// @Success 200 {object} dto.ProblemBody
// @Success 400
// @Success 401
// @Success 404
// @Router /problem/:problem_id/upvote [post]
func (h *Handler) UpvoteProblem(ctx *context.Ctx) error {
	problemId := ctx.Params("problem_id")

	res, err := h.Service.UpvoteProblem(problemId)

	if err != nil {
		return err
	}

	ctx.JSON(res)

	return nil
}

// @Summary Downvote problem
// @Tags Thread
// @Produce json
// @Param problem_id path string true "Problem id"
// @Success 200 {object} dto.ProblemBody
// @Failure 400
// @Success 401
// @Success 404
// @Router /problem/:problem_id/downvote [post]
func (h *Handler) DownvoteProblem(ctx *context.Ctx) error {
	problemId := ctx.Params("problem_id")

	res, err := h.Service.DownvoteProblem(problemId)

	if err != nil {
		return err
	}

	ctx.JSON(res)

	return nil
}

// @Summary Add answer
// @Tags Thread
// @Produce json
// @Param thread_id path string true "Thread id"
// @Param body body dto.AnswerRequestCreateBody true "problem body"
// @Success 200 {object} dto.AnswerBody
// @Failure 400
// @Success 401
// @Success 404
// @Router /thread/:thread_id/answer [post]
func (h *Handler) AddAnswer(ctx *context.Ctx) error {
	userId := ctx.UserId()
	threadId := ctx.Params("thread_id")

	problem := new(dto.AnswerRequestCreateBody)

	if err := ctx.BodyParser(&problem); err != nil {
		return err
	}

	res, err := h.Service.AddAnswer(userId, threadId, problem)

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
