package nlp

import (
	"context"
	"time"

	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/dto"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/proto"
)

type Service struct {
	client proto.NlpClient
}

func (s *Service) Tokenize(body *dto.TokenizeParagraph) (*dto.TokenizeSentences, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.Tokenize(ctx, &proto.TokenizeParagraph{
		Para: body.Para,
	})

	if err != nil {
		return nil, err
	}

	return &dto.TokenizeSentences{
		Sentences: res.Sentences,
	}, nil
}

func NewService(client proto.NlpClient) Service {
	return Service{client}
}
