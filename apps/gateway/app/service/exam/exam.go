package auth

import (
	"context"
	"time"

	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/dto"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/proto"
)

type Service struct {
	client proto.ExamClient
}

func (s *Service) CreateExam(body *dto.ExamCreateBody) (*dto.ExamBody, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.CreateExam(ctx, &proto.ExamCreateBody{
		Year:     body.Year,
		Semester: body.Semester,
		Term:     body.Term,
		CourseId: body.CourseId,
	})

	if err != nil {
		return nil, err
	}

	return proto2dto(res), nil
}

func (s *Service) GetExam(year *int32, semester *string, term *string, courseId *string) ([]*dto.ExamBody, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.GetExams(ctx, &proto.ExamPropertyRequestBody{
		Year:     year,
		Semester: semester,
		Term:     term,
		CourseId: courseId,
	})

	if err != nil {
		return nil, err
	}

	var ret = make([]*dto.ExamBody, 0)

	for _, v := range res.Messages {
		ret = append(ret, proto2dto(v))
	}

	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *Service) UpdateExamByCourseProperty(year int32, semester string, term string, body *dto.ExamUpdateBody) (*dto.ExamBody, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.UpdateExamByCourseProperty(ctx, &proto.ExamRequestUpdateBody{
		Year:     year,
		Semester: semester,
		Term:     term,
		Body: &proto.ExamUpdateBody{
			CourseId:  body.CourseId,
			ThreadIds: body.ThreadIds,
		},
	})

	if err != nil {
		return nil, err
	}

	return proto2dto(res), nil
}

// func (s *Service) DeleteExamByCourseProperty(year int32, semester string, term string) (*dto.ExamBody, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	res, err := s.client.DeleteExamByCourseProperty(ctx, &proto.ExamPropertyRequestBody{
// 		Year:     year,
// 		Semester: semester,
// 		Term:     term,
// 		CourseId: "1233",
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	return proto2dto(res), nil
// }

func proto2dto(p *proto.ExamBody) *dto.ExamBody {
	return &dto.ExamBody{
		Id:        p.XId,
		CourseId:  p.CourseId,
		Year:      p.Year,
		Semester:  p.Semester,
		Term:      p.Term,
		ThreadIds: p.ThreadIds,
	}
}

func NewService(client proto.ExamClient) Service {
	return Service{client}
}
