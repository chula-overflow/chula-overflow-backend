package dto

type TokenizeParagraph struct {
	Para string `json:"para" example:"x + y + z = 3, find x and y"`
}

type TokenizeSentences struct {
	Sentences []string `json:"sentences" example:"find the x, find the y"`
}
