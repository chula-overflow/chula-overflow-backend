package dto

type ThreadRequestCreateBody struct {
	CourseId string `json:"course_id" example:"507f1f77bcf86cd799439011" validate:"required"`
	Year     int32  `json:"year" example:"2011" validate:"required"`
	Semester string `json:"semester" example:"S1" validate:"required"`
	Term     string `json:"term" example:"Final" validate:"required"`
	Question string `json:"question" example:"Final" validate:"required"`
	Answer   string `json:"answer" example:"Final"`
}

type ThreadBody struct {
	Id        string        `json:"_id" example:"507f1f77bcf86cd799439011" validate:"required"`
	ExamId    string        `json:"exam_id" example:"507f1f77bcf86cd799439011" validate:"required"`
	CourseId  string        `json:"course_id" example:"507f1f77bcf86cd799439011" validate:"required"`
	Upvoted   int32         `json:"upvoted" example:"17"`
	Downvoted int32         `json:"downvoted" example:"13"`
	Problems  []ProblemBody `json:"problems"`
	Answers   []AnswerBody  `json:"answers"`
}

type AnswerBody struct {
	Id        string `json:"_id" example:"507f1f77bcf86cd799439011" validate:"required"`
	Body      string `json:"body" example:"I dont know I think it's B" validate:"required"`
	Upvoted   int32  `json:"upvoted" example:"13" validate:"required"`
	Downvoted int32  `json:"downvoted" example:"17" validate:"required"`
}

type ProblemBody struct {
	Id           string `json:"_id" example:"507f1f77bcf86cd799439011" validate:"required"`
	Title        string `json:"title" example:"Hard limit" validate:"required"`
	Body         string `json:"body" example:"What is 'Monad'?" validate:"required"`
	UploadedUser string `json:"uploaded_user" example:"507f1f77bcf86cd799439011" validate:"required"`
	Upvoted      int32  `json:"upvoted" example:"13" validate:"required"`
	Downvoted    int32  `json:"downvoted" example:"17" validate:"required"`
}

type AnswerRequestCreateBody struct {
	Body string `json:"body" example:"What is 'Monad'?" validate:"required"`
}
