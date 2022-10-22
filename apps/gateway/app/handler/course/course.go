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
	GetAllCourseSummary() ([]*dto.CourseSummary, error)
	GetCourse(string) (*dto.Course, error)
}

type Handler struct {
	Service IService
}

// @Summary GetAllCourseSummary
// @Description get detail summary for all course
// @Tags Course
// @Produce json
// @Success 200 {object} []dto.CourseSummary
// @Failure 500
// @Failure 503
// @Router /course [get]
func (h *Handler) GetAllCourseSummary(ctx *context.Ctx) error {
	courses, err := h.Service.GetAllCourseSummary()
	if err != nil {
		if _, ok := status.FromError(err); ok {
			log.Println("Service down")
			return fiber.ErrServiceUnavailable
		} else {
			log.Println("Parse error failed", err)
			return fiber.ErrInternalServerError
		}
	}

	log.Print(courses)
	ctx.JSON(courses)

	return nil
}

// @Summary GetCourse
// @Description Get specific course
// @Tags Course
// @Param course_id path string true "course ID"
// @Success 200 {object} dto.Course
// @Failure 404
// @Failure 500
// @Failure 503
// @Router /course/:course_id [get]
func (h *Handler) GetCourse(ctx *context.Ctx) error {
	courseId := ctx.Params("course_id")

	course, err := h.Service.GetCourse(courseId)

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

	ctx.JSON(course)

	return nil
}

func NewHandler(service IService) Handler {
	return Handler{
		Service: service,
	}
}
