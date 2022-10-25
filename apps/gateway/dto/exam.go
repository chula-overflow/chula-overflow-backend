package dto

type ExamCreateBody struct {
	CourseId string `json:"course_id" example:"507f1f77bcf86cd799439011"`
	Year     int32  `json:"year" example:"2011" validate:"required"`
	Semester string `json:"semester" example:"S1" validate:"required"`
	Term     string `json:"term" example:"Final" validate:"required"`
}

type ExamUpdateBody struct {
	CourseId  string   `json:"course_id" example:"507f1f77bcf86cd799439011"`
	ThreadIds []string `json:"thread_ids" example:"507f1f77bcf86cd799439011,507f1f77bcf86cd799439011"`
}

type ExamBody struct {
	Id        string   `json:"_id" example:"507f1f77bcf86cd799439011" validate:"required"`
	CourseId  string   `json:"course_id" example:"507f1f77bcf86cd799439011" validate:"required"`
	Year      int32    `json:"year" example:"2011" validate:"required"`
	Semester  string   `json:"semester" example:"S1" validate:"required"`
	Term      string   `json:"term" example:"Final" validate:"required"`
	ThreadIds []string `json:"thread_ids" example:"507f1f77bcf86cd799439011,507f1f77bcf86cd799439011" validate:"required"`
}
