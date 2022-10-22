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

func (s *Service) GetAllCourseSummary() ([]*dto.CourseSummary, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.GetAllCourseSummary(ctx, &proto.GetAllCourseSummaryRequest{})

	if err != nil {
		return nil, err
	}

	ret := make([]*dto.CourseSummary, len(res.Courses))

	for i, v := range res.Courses {
		ret[i] = &dto.CourseSummary{
			CourseId:    v.CourseId,
			CourseName:  v.CourseName,
			ThreadCount: v.ThreadCount,
		}
	}

	return ret, nil
}

func (s *Service) GetCourse(courseId string) (*dto.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.GetCourse(ctx, &proto.GetCourseRequest{
		CourseId: courseId,
	})

	if err != nil {
		return nil, err
	}

	exams := make([]dto.ExamSummary, len(res.Exams))

	for i, v := range res.Exams {
		exams[i] = dto.ExamSummary{
			ExamId:      v.ExamId,
			ExamName:    v.ExamName,
			ThreadCount: v.ThreadCount,
		}
	}

	return &dto.Course{
		CourseName: res.CourseName,
		Exams:      exams,
	}, nil
}

func NewService(client proto.CourseClient) Service {
	return Service{
		client: client,
	}
}
