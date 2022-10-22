package auth

import (
	"context"
	"time"

	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/dto"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/proto"
)

type Service struct {
	client proto.AuthClient
}

func (s *Service) Login(email string) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.Login(ctx, &proto.AuthRequest{
		Email: email,
	})

	if err != nil {
		return nil, err
	}

	return &res.Token, nil
}

func (s *Service) Revoke(sessionId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.client.Revoke(ctx, &proto.RevokeRequest{
		Token: sessionId,
	})

	return err
}

func (s *Service) Me(sessionId string) (*dto.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := s.client.Me(ctx, &proto.MeRequest{
		Token: sessionId,
	})

	if err != nil {
		return nil, err
	}

	return &dto.User{
		Email: user.Email,
	}, nil
}

func NewService(client proto.AuthClient) Service {
	return Service{
		client: client,
	}
}
