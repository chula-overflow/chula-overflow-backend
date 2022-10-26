package auth

import (
	"strconv"

	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/validator"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/context"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/dto"
	"github.com/gofiber/fiber/v2"
)

type IService interface {
	CreateExam(*dto.ExamCreateBody) (*dto.ExamBody, error)
	GetExam(year *int32, semester *string, term *string, courseId *string) ([]*dto.ExamBody, error)
	UpdateExamByCourseProperty(year int32, semester string, term string, body *dto.ExamUpdateBody) (*dto.ExamBody, error)
	// DeleteExamByCourseProperty(year int32, semester string, term string) (*dto.ExamBody, error)
}

type Handler struct {
	Service IService
	v       *validator.MyValidator
}

// @Summary Get all exam
// @Description Choose 1 type of these 3 types:
// @Description
// @Description Get all (No query)
// @Description
// @Param year query int false "year"
// @Param semester query string false "semester"
// @Param term query string false "term"
// @Param course_id query string false "Course id"
// @Tags Exam
// @Produce json
// @Success 200 {object} []dto.ExamBody
// @Failure 400
// @Router /exam [get]
func (h *Handler) GetExam(ctx *context.Ctx) error {
	year := ctx.Query("year")
	semester := ctx.Query("semester")
	term := ctx.Query("term")
	courseId := ctx.Query("course_id")

	yearInt, err := strconv.Atoi(year)
	var yearRef *int32
	if err != nil {
		yearRef = nil
	} else {
		year32 := int32(yearInt)
		yearRef = &year32
	}

	res, err := h.Service.GetExam(yearRef, &semester, &term, &courseId)

	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

// @Summary Create exam
// @Tags Exam
// @Accept json
// @Produce json
// @Param exam_body body dto.ExamCreateBody true "Exam body"
// @Success 201 {object} dto.ExamBody
// @Failure 400
// @Failure 401
// @Failure 422
// @Router /exam [post]
func (h *Handler) CreateExam(ctx *context.Ctx) error {
	body := new(dto.ExamCreateBody)

	if err := ctx.BodyParser(body); err != nil {
		return err
	}
	err := h.v.Struct(body)
	if err != nil {
		return err
	}

	res, err := h.Service.CreateExam(body)

	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

// @Summary Update exam
// @Tags Exam
// @Produce json
// @Param update_body body dto.ExamUpdateBody true "Update body"
// @Param year query int true "year"
// @Param semester query string true "semester"
// @Param term query string true "term"
// @Success 204 {object} dto.ExamBody
// @Failure 400
// @Failure 401
// @Failure 404
// @Router /exam [put]
func (h *Handler) UpdateExamByCourseProperty(ctx *context.Ctx) error {
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

	body := new(dto.ExamUpdateBody)
	if err := ctx.BodyParser(body); err != nil {
		return err
	}

	res, err := h.Service.UpdateExamByCourseProperty(int32(yearInt), semester, term, body)

	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

// @Summary Delete exam
// @Tags Exam
// @Param year query int true "year"
// @Param semester query string true "semester"
// @Param term query string true "term"
// @Produce json
// @Success 204 {object} dto.ExamBody
// @Failure 400
// @Failure 401
// @Failure 404
// @Router /exam [delete]
// func (h *Handler) DeleteExam(ctx *context.Ctx) error {
// 	year := ctx.Query("year")
// 	semester := ctx.Query("semester")
// 	term := ctx.Query("term")

// 	if year == "" || semester == "" || term == "" {
// 		return fiber.ErrBadRequest
// 	}

// 	yearInt, err := strconv.Atoi(year)

// 	if err != nil {
// 		return fiber.ErrBadRequest
// 	}

// 	res, err := h.Service.DeleteExamByCourseProperty(int32(yearInt), semester, term)

// 	if err != nil {
// 		return err
// 	}

// 	return ctx.JSON(res)
// }

func NewHandler(service IService, validator *validator.MyValidator) Handler {
	return Handler{service, validator}
}
