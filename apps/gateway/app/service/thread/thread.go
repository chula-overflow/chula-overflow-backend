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

func (s *Service) GetThread(threadId string) (*dto.Thread, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.GetThread(ctx, &proto.GetThreadRequest{
		ThreadId: threadId,
	})

	if err != nil {
		return nil, err
	}

	replies := make([]dto.Reply, len(res.Replies))

	for i, v := range res.Replies {
		replies[i] = dto.Reply{
			ThreadId: v.ThreadId,
			Body:     v.Body,
			Upvote:   v.Upvote,
			Downvote: v.Downvote,
		}
	}

	return &dto.Thread{
		ThreadId:           threadId,
		ProblemName:        res.ProblemName,
		ProblemDescription: res.ProblemDescription,
		Upvote:             res.Upvote,
		Downvote:           res.Downvote,
		Tags:               res.Tags,
		Replies:            replies,
	}, nil
}

func (s *Service) CreateThread(createThread *dto.CreateThread) (*dto.CreateThreadResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.CreateThread(ctx, &proto.CreateThreadRequest{
		CourseName:         createThread.CourseName,
		ExamName:           createThread.ExamName,
		ProblemDescription: createThread.ProblemDescription,
		ProblemName:        createThread.ProblemName,
		Tags:               createThread.Tags,
	})

	if err != nil {
		return nil, err
	}

	return &dto.CreateThreadResponse{
		ThreadId: res.ThreadId,
	}, nil
}

func (s *Service) CreateReply(createReply *dto.CreateReply, threadId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.client.CreateReply(ctx, &proto.CreateReplyRequest{
		Body:     createReply.Body,
		ThreadId: threadId,
	})

	return err
}

func NewService(client proto.ThreadClient) Service {
	return Service{
		client: client,
	}
}
