package nlp

import (
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/app/validator"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/context"
	"github.com/chula-overflow/chula-overflow-backend/apps/gateway/dto"
)

type IService interface {
	Tokenize(*dto.TokenizeParagraph) (*dto.TokenizeSentences, error)
}

type Handler struct {
	Service IService
	v       *validator.MyValidator
}

// @Summary Tokenize
// @Param paragraph body dto.TokenizeParagraph true "Paragraph"
// @Tags NLP
// @Produce json
// @Success 200 {object} dto.TokenizeSentences
// @Failure 400
// @Router /tokenize [post]
func (h *Handler) Tokenize(ctx *context.Ctx) error {
	para := new(dto.TokenizeParagraph)

	if err := ctx.BodyParser(para); err != nil {
		return err
	}

	res, err := h.Service.Tokenize(para)

	if err != nil {
		return err
	}

	ctx.JSON(res)

	return nil
}

func NewHandler(service IService, validator *validator.MyValidator) Handler {
	return Handler{service, validator}
}
