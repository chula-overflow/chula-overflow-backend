package dto

type CourseBody struct {
	Id             string   `json:"_id" example:"507f1f77bcf86cd799439011" validate:"required"`
	CourseId       string   `json:"course_id" example:"507f1f77bcf86cd799439011" validate:"required"`
	CourseName     string   `json:"course_name" example:"Calculus" validate:"required"`
	CourseCodename string   `json:"course_codename" example:"2112101" validate:"required"`
	ExamIds        []string `json:"exam_ids" example:"507f1f77bcf86cd799439011,507f1f77bcf86cd799439011"`
}

type CourseCreateBody struct {
	CourseId       string `json:"course_id" example:"507f1f77bcf86cd799439011" validate:"required"`
	CourseName     string `json:"course_name" example:"Calculus" validate:"required"`
	CourseCodename string `json:"course_codename" example:"2112101" validate:"required"`
}

type CourseUpdateBody struct {
	CourseName     string   `json:"course_name" example:"Calculus"`
	CourseCodename string   `json:"course_codename" example:"2112101"`
	ExamIds        []string `json:"exam_ids" example:"507f1f77bcf86cd799439011,507f1f77bcf86cd799439011"`
}
