package auth

import (
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/validator"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/context"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/dto"
)

type IService interface {
	CreateCourse(*dto.CourseCreateBody) (*dto.CourseBody, error)
	GetAllCourses() ([]*dto.CourseBody, error)
	GetCourseByCourseId(string) (*dto.CourseBody, error)
	UpdateCourse(string, *dto.CourseUpdateBody) (*dto.CourseBody, error)
	DeleteCourse(string) (*dto.CourseBody, error)
}

type Handler struct {
	Service IService
	v       *validator.MyValidator
}

// @Summary Get all course
// @Tags Course
// @Produce json
// @Success 200 {object} []dto.CourseBody
// @Router /course [get]
func (h *Handler) GetAllCourses(ctx *context.Ctx) error {
	res, err := h.Service.GetAllCourses()

	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

// @Summary Get course
// @Tags Course
// @Produce json
// @Param course_id path string true "Course id"
// @Success 200 {object} dto.CourseBody
// @Failure 404
// @Router /course/{course_id} [get]
func (h *Handler) GetCourseByCourseId(ctx *context.Ctx) error {
	courseId := ctx.Params("course_id")

	res, err := h.Service.GetCourseByCourseId(courseId)

	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

// @Summary Create Course
// @Tags Course
// @Accept json
// @Produce json
// @Param course_body body dto.CourseCreateBody true "Course body"
// @Success 201 {object} dto.CourseBody
// @Failure 400
// @Failure 401
// @Failure 422
// @Router /course [post]
func (h *Handler) CreateCourse(ctx *context.Ctx) error {
	body := new(dto.CourseCreateBody)

	if err := ctx.BodyParser(body); err != nil {
		return err
	}
	err := h.v.Struct(body)
	if err != nil {
		return err
	}

	res, err := h.Service.CreateCourse(body)

	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

// @Summary Update course
// @Tags Course
// @Produce json
// @Param body body dto.CourseUpdateBody true "Update body"
// @Param course_id path string true "Course id"
// @Success 204 {object} dto.CourseBody
// @Failure 401
// @Failure 404
// @Router /course/{course_id} [put]
func (h *Handler) UpdateCourse(ctx *context.Ctx) error {
	courseId := ctx.Params("course_id")

	body := new(dto.CourseUpdateBody)

	if err := ctx.BodyParser(body); err != nil {
		return err
	}
	err := h.v.Struct(body)
	if err != nil {
		return err
	}

	res, err := h.Service.UpdateCourse(courseId, body)

	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

// @Summary Delete course
// @Tags Course
// @Produce json
// @Param Course_id path string true "Course id"
// @Success 204 {object} dto.CourseBody
// @Failure 401
// @Failure 404
// @Router /course/{course_id} [delete]
func (h *Handler) DeleteCourse(ctx *context.Ctx) error {
	courseId := ctx.Params("course_id")

	res, err := h.Service.DeleteCourse(courseId)

	if err != nil {
		return err
	}

	return ctx.JSON(res)
}

func NewHandler(service IService, validator *validator.MyValidator) Handler {
	return Handler{service, validator}
}
