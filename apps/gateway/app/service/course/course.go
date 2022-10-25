package auth

import (
	"context"
	"time"

	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/dto"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/proto"
)

type Service struct {
	client proto.CourseClient
}

func (s *Service) CreateCourse(body *dto.CourseCreateBody) (*dto.CourseBody, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.CreateCourse(ctx, &proto.CourseCreateBody{
		CourseId:       body.CourseId,
		CourseName:     body.CourseName,
		CourseCodename: body.CourseCodename,
	})

	if err != nil {
		return nil, err
	}

	return proto2dto(res), nil
}

func (s *Service) GetAllCourses() ([]*dto.CourseBody, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.GetAllCourses(ctx, &proto.GetAllCourseBody{})

	if err != nil {
		return nil, err
	}

	var ret = make([]*dto.CourseBody, 0)

	for _, v := range res.Messages {
		ret = append(ret, proto2dto(v))
	}

	if err != nil {
		return nil, err
	}

	return ret, nil
}
func (s *Service) GetCourseByCourseId(courseId string) (*dto.CourseBody, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.GetCourseByCourseId(ctx, &proto.GetCourseRequest{
		CourseId: courseId,
	})

	if err != nil {
		return nil, err
	}

	return proto2dto(res), nil
}
func (s *Service) UpdateCourse(courseId string, body *dto.CourseUpdateBody) (*dto.CourseBody, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.UpdateCourse(ctx, &proto.CourseRequestUpdateBody{
		CourseId: courseId,
		Body: &proto.CourseUpdateBody{
			CourseId:       &courseId,
			CourseName:     &body.CourseName,
			CourseCodename: &body.CourseCodename,
			ExamIds:        body.ExamIds,
		},
	})

	if err != nil {
		return nil, err
	}

	return proto2dto(res), nil
}

func (s *Service) DeleteCourse(courseId string) (*dto.CourseBody, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.DeleteCourse(ctx, &proto.GetCourseRequest{
		CourseId: courseId,
	})

	if err != nil {
		return nil, err
	}

	return proto2dto(res), nil
}

func proto2dto(p *proto.CourseBody) *dto.CourseBody {
	return &dto.CourseBody{
		Id:             p.XId,
		CourseId:       p.CourseId,
		CourseName:     p.CourseName,
		CourseCodename: p.CourseCodename,
		ExamIds:        p.ExamIds,
	}
}

func NewService(client proto.CourseClient) Service {
	return Service{client}
}
