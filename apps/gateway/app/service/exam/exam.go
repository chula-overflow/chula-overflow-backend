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

func (s *Service) GetExam(examId string) (*dto.Exam, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.GetExam(ctx, &proto.GetExamRequest{
		ExamId: examId,
	})

	if err != nil {
		return nil, err
	}

	threads := make([]dto.ThreadSummary, len(res.Threads))

	for i, v := range res.Threads {
		threads[i] = dto.ThreadSummary{
			ThreadId:           v.ThreadId,
			ProblemName:        v.ProblemName,
			ProblemDescription: v.ProblemDescription,
			Upvote:             v.Upvote,
			Downvote:           v.Downvote,
			Tags:               v.Tags,
		}
	}

	return &dto.Exam{
		ExamName: res.ExamName,
		Threads:  threads,
	}, nil
}

func NewService(client proto.ExamClient) Service {
	return Service{
		client: client,
	}
}
