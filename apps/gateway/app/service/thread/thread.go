package auth

import (
	"context"
	"time"

	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/dto"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/proto"
)

type Service struct {
	client proto.ThreadClient
}

func (s *Service) DownvoteThread(threadId string) (*dto.ThreadBody, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.DownvoteThread(ctx, &proto.ThreadIdRequestBody{
		ThreadId: threadId,
	})

	if err != nil {
		return nil, err
	}

	return proto2dto(res), nil
}

func (s *Service) UpvoteThread(threadId string) (*dto.ThreadBody, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.UpvoteThread(ctx, &proto.ThreadIdRequestBody{
		ThreadId: threadId,
	})

	if err != nil {
		return nil, err
	}

	return proto2dto(res), nil
}

func (s *Service) GetThreadById(threadId string) (*dto.ThreadBody, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.GetThreadById(ctx, &proto.ThreadIdRequestBody{
		ThreadId: threadId,
	})

	if err != nil {
		return nil, err
	}

	return proto2dto(res), nil
}

func (s *Service) GetAllThreadsByExamProperty(year int32, semester string, term string) ([]*dto.ThreadBody, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.GetAllThreadsByExamProperty(ctx, &proto.ExamPropertyRequestBody{
		Year:     year,
		Semester: semester,
		Term:     term,
	})

	if err != nil {
		return nil, err
	}

	var ret = make([]*dto.ThreadBody, 0)

	for _, v := range res.Messages {
		ret = append(ret, proto2dto(v))
	}

	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *Service) CreateThread(userId string, body *dto.ThreadRequestCreateBody) (*dto.ThreadBody, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.CreateThread(ctx, &proto.ThreadRequestCreateBody{
		CourseId:     body.CourseId,
		Year:         body.Year,
		Semester:     body.Semester,
		Term:         body.Term,
		UploadedUser: userId,
		Question:     body.Question,
		Answer:       &body.Answer,
	})

	if err != nil {
		return nil, err
	}

	return proto2dto(res), nil
}

func NewService(client proto.ThreadClient) Service {
	return Service{
		client: client,
	}
}

func proto2dto(p *proto.ThreadBody) *dto.ThreadBody {
	problems := make([]dto.Problem, 0)

	for _, v := range p.Problems {
		problems = append(problems, dto.Problem{
			Title:        *v.Title,
			Body:         *v.Body,
			UploadedUser: *v.UploadedUser,
			Upvoted:      *v.Upvoted,
			Downvoted:    *v.Downvoted,
		})
	}

	answers := make([]dto.Answer, 0)

	for _, v := range p.Answers {
		answers = append(answers, dto.Answer{
			Body:      *v.Body,
			Upvoted:   *v.Upvoted,
			Downvoted: *v.Downvoted,
		})
	}

	return &dto.ThreadBody{
		Id:        p.XId,
		ExamId:    p.ExamId,
		CourseId:  p.CourseId,
		Upvoted:   p.Upvoted,
		Downvoted: p.Downvoted,
		Problems:  problems,
		Answers:   answers,
	}
}
