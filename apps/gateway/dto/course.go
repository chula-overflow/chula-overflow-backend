package dto

type CourseBody struct {
	Id             string   `json:"_id" example:"507f1f77bcf86cd799439011" validate:"required"`
	CourseId       string   `json:"course_id" example:"2100111" validate:"required"`
	CourseName     string   `json:"course_name" example:"CALCULUS 1" validate:"required"`
	CourseCodename string   `json:"course_codename" example:"CAL" validate:"required"`
	ExamIds        []string `json:"exam_ids" example:"507f1f77bcf86cd799439011,507f1f77bcf86cd799439011"`
}

type CourseCreateBody struct {
	CourseId       string `json:"course_id" example:"2100111" validate:"required"`
	CourseName     string `json:"course_name" example:"CALCULUS 1" validate:"required"`
	CourseCodename string `json:"course_codename" example:"CAL" validate:"required"`
}

type CourseUpdateBody struct {
	CourseName     string   `json:"course_name" example:"CALCULUS 1"`
	CourseCodename string   `json:"course_codename" example:"CAL"`
	ExamIds        []string `json:"exam_ids" example:"507f1f77bcf86cd799439011,507f1f77bcf86cd799439011"`
}
