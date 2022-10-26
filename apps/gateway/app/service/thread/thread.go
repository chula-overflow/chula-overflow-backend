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

func (s *Service) GetThreadById(threadId string) (*dto.ThreadBody, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.GetThreadById(ctx, &proto.ThreadIdRequestBody{
		ThreadId: threadId,
	})

	if err != nil {
		return nil, err
	}

	return threadProto2Dto(res), nil
}

func (s *Service) GetThread(year int32, semester string, term string, courseId string) ([]*dto.ThreadBody, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.GetThread(ctx, &proto.ThreadPropertyRequestBody{
		Year:     &year,
		Semester: &semester,
		Term:     &term,
		CourseId: &courseId,
	})

	if err != nil {
		return nil, err
	}

	var ret = make([]*dto.ThreadBody, 0)

	for _, v := range res.Messages {
		ret = append(ret, threadProto2Dto(v))
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

	return threadProto2Dto(res), nil
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

	return threadProto2Dto(res), nil
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

	return threadProto2Dto(res), nil
}

func (s *Service) DownvoteAnswer(answerId string) (*dto.AnswerBody, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.DownvoteAnswer(ctx, &proto.ThreadAnswerIdRequestBody{
		AnswerId: answerId,
	})

	if err != nil {
		return nil, err
	}

	return ansProto2Dto(res), nil
}

func (s *Service) UpvoteAnswer(answerId string) (*dto.AnswerBody, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.UpvoteAnswer(ctx, &proto.ThreadAnswerIdRequestBody{
		AnswerId: answerId,
	})

	if err != nil {
		return nil, err
	}

	return ansProto2Dto(res), nil
}

func (s *Service) DownvoteProblem(problemId string) (*dto.ProblemBody, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.DownvoteProblem(ctx, &proto.ThreadProblemIdRequestBody{
		ProblemId: problemId,
	})

	if err != nil {
		return nil, err
	}

	return problemProto2Dto(res), nil
}

func (s *Service) UpvoteProblem(problemId string) (*dto.ProblemBody, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.UpvoteProblem(ctx, &proto.ThreadProblemIdRequestBody{
		ProblemId: problemId,
	})

	if err != nil {
		return nil, err
	}

	return problemProto2Dto(res), nil
}

func (s *Service) AddAnswer(userId string, threadId string, body *dto.AnswerRequestCreateBody) (*dto.AnswerBody, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.AddAnswer(ctx, &proto.AnswerRequestCreateBody{
		ThreadId: threadId,
		Body:     body.Body,
		UserId:   userId,
	})

	if err != nil {
		return nil, err
	}

	return ansProto2Dto(res), nil
}

func threadProto2Dto(p *proto.ThreadBody) *dto.ThreadBody {
	problems := make([]dto.ProblemBody, 0)

	for _, v := range p.Problems {
		problems = append(problems, *problemProto2Dto(v))
	}

	answers := make([]dto.AnswerBody, 0)

	for _, v := range p.Answers {
		answers = append(answers, *ansProto2Dto(v))
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

func ansProto2Dto(p *proto.AnswerBody) *dto.AnswerBody {
	return &dto.AnswerBody{
		Id:        p.Id,
		Body:      p.Body,
		Upvoted:   p.Upvoted,
		Downvoted: p.Downvoted,
	}
}

func problemProto2Dto(p *proto.ProblemBody) *dto.ProblemBody {
	return &dto.ProblemBody{
		Id:           p.Id,
		Title:        p.Title,
		Body:         p.Body,
		UploadedUser: p.UploadedUser,
		Upvoted:      p.Upvoted,
		Downvoted:    p.Downvoted,
	}
}

func NewService(client proto.ThreadClient) Service {
	return Service{
		client: client,
	}
}
